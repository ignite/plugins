package ssh

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ignite/cli/v28/ignite/pkg/errors"
	"github.com/ignite/cli/v28/ignite/pkg/gocmd"
	"github.com/ignite/cli/v28/ignite/pkg/randstr"
	"github.com/melbahja/goph"
	"github.com/pkg/sftp"
	"golang.org/x/sync/errgroup"
)

const (
	workdir = "spaceship"
)

type SSH struct {
	username    string
	password    string
	host        string
	port        string
	rawKey      string
	key         string
	keyPassword string
	workspace   string
	client      *goph.Client
	sftpClient  *sftp.Client
}

// Option configures ssh config.
type Option func(*SSH) error

// WithUser set SSH username.
func WithUser(username string) Option {
	return func(o *SSH) error {
		if o.username != "" {
			return nil
		}
		o.username = strings.TrimSpace(username)
		return nil
	}
}

// WithPassword set SSH password.
func WithPassword(password string) Option {
	return func(o *SSH) error {
		if o.password != "" {
			return nil
		}
		o.password = strings.TrimSpace(password)
		return nil
	}
}

// WithPort set SSH port.
func WithPort(port string) Option {
	return func(o *SSH) error {
		if o.port != "" {
			return nil
		}
		o.port = strings.TrimSpace(port)
		return nil
	}
}

// WithRawKey set SSH raw key.
func WithRawKey(rawKey string) Option {
	return func(o *SSH) error {
		if o.rawKey != "" {
			return nil
		}
		o.rawKey = strings.TrimSpace(rawKey)
		return nil
	}
}

// WithKey set SSH key.
func WithKey(key string) Option {
	return func(o *SSH) error {
		if o.key != "" {
			return nil
		}
		o.key = strings.TrimSpace(key)
		return nil
	}
}

// WithKeyPassword set SSH key password.
func WithKeyPassword(keyPassword string) Option {
	return func(o *SSH) error {
		if o.keyPassword != "" {
			return nil
		}
		o.keyPassword = strings.TrimSpace(keyPassword)
		return nil
	}
}

// WithWorkspace set SSH workspace.
func WithWorkspace(workspace string) Option {
	return func(o *SSH) error {
		o.workspace = strings.TrimSpace(workspace)
		return nil
	}
}

// New creates a new ssh object.
func New(host string, options ...Option) (*SSH, error) {
	host, port, username, password, err := parseURI(host)
	if err != nil {
		return nil, err
	}
	s := &SSH{
		username:  username,
		host:      host,
		port:      port,
		password:  password,
		workspace: randstr.Runes(10),
	}
	for _, apply := range options {
		if err := apply(s); err != nil {
			return nil, err
		}
	}
	return s, s.validate()
}

func parseURI(uri string) (host string, port string, username string, password string, err error) {
	uri = strings.TrimSpace(uri)
	if !strings.HasPrefix(uri, "ssh://") {
		uri = "ssh://" + uri
	}
	parsedURL, err := url.Parse(uri)
	if err != nil {
		return "", "", "", "", errors.Wrapf(err, "error parsing URI %s", uri)
	}
	host = parsedURL.Hostname()
	port = parsedURL.Port()
	if port == "" {
		port = "22"
	}

	if parsedURL.User != nil {
		username = parsedURL.User.Username()
		password, _ = parsedURL.User.Password()
	}
	if username == "" {
		username = "root"
	}
	return host, port, username, password, nil
}

func (s *SSH) Workspace() string {
	return filepath.Join(workdir, s.workspace)
}

func (s *SSH) Bin() string {
	return filepath.Join(s.Workspace(), "bin")
}

func (s *SSH) Home() string {
	return filepath.Join(s.Workspace(), "home")
}

func (s *SSH) Genesis() string {
	return filepath.Join(s.Home(), "config", "genesis.json")
}

func (s *SSH) Log() string {
	return filepath.Join(s.Workspace(), "log")
}

func (s *SSH) RunnerScript() string {
	return filepath.Join(s.Workspace(), "run.sh")
}

func (s *SSH) validate() error {
	switch {
	case s.username == "":
		return fmt.Errorf("ssh username is required")
	case s.key != "" && s.rawKey != "":
		return errors.New("ssh key and raw key are both set")
	case s.key != "" && s.password != "":
		return errors.New("ssh key and password are both set")
	case s.rawKey != "" && s.password != "":
		return errors.New("ssh raw key and password are both set")
	default:
		return nil
	}
}

func (s *SSH) auth() (goph.Auth, error) {
	switch {
	case s.rawKey != "":
		return goph.RawKey(s.rawKey, s.keyPassword)
	case s.key != "":
		return goph.Key(s.key, s.keyPassword)
	case s.password != "":
		return goph.Password(s.password), nil
	default:
		return goph.KeyboardInteractive(s.password), nil
	}
}

// Close closes the SSH client.
func (s *SSH) Close() error {
	if err := s.sftpClient.Close(); err != nil {
		return err
	}
	return s.client.Close()
}

// Connect connects the SSH client.
func (s *SSH) Connect() error {
	auth, err := s.auth()
	if err != nil {
		return err
	}

	s.client, err = goph.New(s.username, s.host, auth)
	if err != nil {
		return errors.Wrapf(err, "Failed to connect to %v", s)
	}

	s.sftpClient, err = s.client.NewSftp()
	if err != nil {
		return err
	}

	return s.ensureEnvironment()
}

func (s *SSH) ensureEnvironment() error {
	if err := s.sftpClient.MkdirAll(s.Bin()); err != nil {
		return errors.Wrapf(err, "failed to create bin dir %s", s.Bin())
	}
	if err := s.sftpClient.MkdirAll(s.Home()); err != nil {
		return errors.Wrapf(err, "failed to create home dir %s", s.Home())
	}
	if err := s.sftpClient.MkdirAll(s.Log()); err != nil {
		return errors.Wrapf(err, "failed to create home dir %s", s.Home())
	}
	return nil
}

func (s *SSH) ensureLocalBin(name string) error {
	// find ignite binary path
	path, err := exec.LookPath(name)
	if err != nil {
		return err
	}
	_, err = s.UploadBinary(path)
	if err != nil {
		return err
	}
	return nil
}

func (s *SSH) Upload(ctx context.Context, srcPath, dstPath string) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.SetLimit(5)
	err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			rel, err := filepath.Rel(srcPath, path)
			if err != nil {
				return err
			}
			// skip hidden files and folders.
			if strings.HasPrefix(rel, ".") {
				return nil
			}
			newPath := filepath.Join(dstPath, rel)

			grp.Go(func() error {
				if err := s.UploadFile(path, newPath); err != nil {
					return errors.Wrapf(err, "failed to upload file %s to %s", path, newPath)
				}
				return nil
			})
		}
		return nil
	})
	if err != nil {
		return err
	}
	return grp.Wait()
}

func (s *SSH) UploadFile(filePath, dstPath string) error {
	dstDir := filepath.Dir(dstPath)
	if err := s.sftpClient.MkdirAll(dstDir); err != nil {
		return errors.Wrapf(err, "failed to create destiny path %s", dstDir)
	}

	srcPath, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	return s.client.Upload(srcPath, dstPath)
}

func (s *SSH) UploadBinary(srcPath string) (string, error) {
	var (
		filename = filepath.Base(srcPath)
		binPath  = filepath.Join(s.Bin(), filename)
	)
	if err := s.UploadFile(srcPath, binPath); err != nil {
		return "", err
	}

	// give binary permission
	if err := s.sftpClient.Chmod(binPath, 0o755); err != nil {
		return "", err
	}
	return binPath, nil
}

func (s *SSH) UploadRunnerScript(srcPath string) (string, error) {
	path := s.RunnerScript()
	if err := s.UploadFile(srcPath, s.RunnerScript()); err != nil {
		return "", err
	}

	// give binary permission
	if err := s.sftpClient.Chmod(path, 0o755); err != nil {
		return "", err
	}
	return path, nil
}

func (s *SSH) UploadHome(ctx context.Context, srcPath string) (string, error) {
	path := s.Home()
	return path, s.Upload(ctx, srcPath, path)
}

func (s *SSH) RunCommand(ctx context.Context, name string, args ...string) (string, error) {
	cmd, err := s.client.CommandContext(ctx, name, args...)
	if err != nil {
		return "", err
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func (s *SSH) Start(ctx context.Context) (string, error) {
	return s.runScript(ctx, "start")
}

func (s *SSH) Restart(ctx context.Context) (string, error) {
	return s.runScript(ctx, "restart")
}

func (s *SSH) Stop(ctx context.Context) (string, error) {
	return s.runScript(ctx, "stop")
}

func (s *SSH) Status(ctx context.Context) (string, error) {
	return s.runScript(ctx, "status")
}

func (s *SSH) runScript(ctx context.Context, args ...string) (string, error) {
	return s.RunCommand(ctx, s.RunnerScript(), args...)
}

func (s *SSH) HasInitialized(ctx context.Context) bool {
	return s.Exist(ctx, s.Genesis())
}

func (s *SSH) Exist(ctx context.Context, path string) bool {
	cmd := fmt.Sprintf("[ -f '%s' ] && echo 'true'", path)
	exist, err := s.RunCommand(ctx, cmd)
	if err != nil {
		return false
	}
	return exist == "true"
}

func (s *SSH) OS(ctx context.Context) (string, error) {
	v, err := s.Uname(ctx)
	if err != nil {
		return "", err
	}
	return strings.ToLower(v), nil
}

func (s *SSH) Arch(ctx context.Context) (string, error) {
	v, err := s.Uname(ctx, "-m")
	if err != nil {
		return "", err
	}
	return strings.ToLower(v), nil
}

func (s *SSH) Target(ctx context.Context) (string, error) {
	os, err := s.OS(ctx)
	if err != nil {
		return "", err
	}

	arch, err := s.Arch(ctx)
	if err != nil {
		return "", err
	}

	return gocmd.BuildTarget(os, arch), nil
}

func (s *SSH) Uname(ctx context.Context, args ...string) (string, error) {
	return s.RunCommand(ctx, "uname", args...)
}