/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package fs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/pkg/errors"
	"net/http"
)

const (
	VarDir         = "var"
	WwwDir         = "www"
	AssetsDir      = "assets"
	DbDir          = "db"
	KTDir          = "kt"
	ProjectDir     = "project"
	ModelDir       = "model"
	LibDir         = "lib"
	OutDir         = "out"
	TmpDir         = "tmp"
	LogDir         = "log"
	DirPerm        = 0755
	FilePerm       = 0666
	KTPerm         = 0600
	PackExt        = ".steam"
	KindEngine     = "engine"
	KindFile       = "file"
	KindExperiment = "module"
)

func NewID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("Failed generating ID: %v", err)
	}
	return u.String(), nil
}

func ResolvePath(p string) (string, error) {
	if !path.IsAbs(p) {
		wd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("Failed identifying working directory: %v", err)
		}
		return path.Join(wd, p), nil
	}
	return p, nil
}

func MkWorkingDirectory(p string) (string, error) {
	wd, err := ResolvePath(p)
	if err != nil {
		return "", err
	}

	dirs := []string{DbDir, ProjectDir, ModelDir, LibDir, TmpDir, LogDir}

	for _, dir := range dirs {
		if err := os.MkdirAll(path.Join(wd, dir), DirPerm); err != nil {
			return "", err
		}
	}

	return wd, nil
}

var nameRegexp = regexp.MustCompile(`(?i)^[a-z0-9][a-z0-9._-]{0,127}$`)

func ValidateName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("Name cannot be empty")
	}
	if !nameRegexp.MatchString(name) {
		return fmt.Errorf("Name must match regexp %s", nameRegexp.String())
	}
	return nil
}

func FileExists(p string) bool {
	s, err := os.Stat(p)
	if err != nil {
		return false
	}

	if s.IsDir() {
		return false
	}

	return true
}

func DirExists(p string) bool {
	s, err := os.Stat(p)
	if err != nil {
		return false
	}

	if !s.IsDir() {
		return false
	}

	return true
}

func Mkdir(p string) error {
	return os.MkdirAll(p, DirPerm)
}

func Rmdir(p string) error {
	return os.RemoveAll(p)
}

func Rm(p string) error {
	return os.Remove(p)
}

func ListDirs(p string) ([]string, error) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			names = append(names, file.Name())
		}
	}
	return names, nil
}

func ListFiles(p string) ([]string, error) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0)
	for _, file := range files {
		if !file.IsDir() {
			name := file.Name()
			if name != ".steam" {
				names = append(names, name)
			}
		}
	}
	return names, nil
}

func GetPack(wd, kind, pack string) (string, bool) {
	p := GetPackPath(wd, kind, pack)

	if !FileExists(p) {
		return p, false
	}

	return p, true
}

type Package struct {
	Name       string
	ModifiedAt time.Time
}

func GetPacks(wd, kind string) ([]*Package, error) {
	packDir := path.Join(wd, LibDir, kind)
	files, err := ioutil.ReadDir(packDir)
	if err != nil {
		return nil, fmt.Errorf("Pack directory read failed: %s: %v", packDir, err)
	}

	extLen := len(PackExt)
	packs := make([]*Package, 0)
	for _, file := range files {
		if !file.IsDir() {
			name := file.Name()
			if path.Ext(name) == PackExt {
				packs = append(packs, &Package{name[:len(name)-extLen], file.ModTime()})
			}
		}
	}

	return packs, nil
}

func GetWwwRoot(wd string) string {
	return path.Join(wd, WwwDir)
}

func GetProjectPath(wd string, projectId int64) string {
	location := strconv.FormatInt(projectId, 10)
	return path.Join(wd, ProjectDir, location)
}

func GetPackagePath(wd string, projectId int64, packageName string) string {
	return path.Join(GetProjectPath(wd, projectId), packageName)
}

func GetPackageRelativePath(wd string, projectId int64, packageName, relativePath string) (string, error) {
	packagePath := GetPackagePath(wd, projectId, packageName)
	absPath := path.Join(packagePath, relativePath)
	if !strings.HasPrefix(absPath, packagePath) {
		return "", fmt.Errorf("Relative path %s is not contained in the package %s", relativePath, packageName)
	}
	return absPath, nil
}

func MapToJson(attrs map[string]string) ([]byte, error) {
	b, err := json.Marshal(attrs)
	if err != nil {
		return nil, fmt.Errorf("Failed encoding attributes: %s", err)
	}
	return b, nil
}

func JsonToMap(b []byte) (map[string]string, error) {
	decoder := json.NewDecoder(bytes.NewBuffer(b))
	var attrs map[string]string
	for {
		if err := decoder.Decode(&attrs); err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("Failed decoding attributes: %s", err)
		}
	}
	return attrs, nil
}

func GetPackageAttributes(wd string, projectId int64, packageName string) ([]byte, error) {
	b, err := ioutil.ReadFile(path.Join(GetPackagePath(wd, projectId, packageName), ".steam"))
	if err != nil {
		return nil, fmt.Errorf("Failed reading attributes: %s", err)
	}
	return b, nil
}

func SetPackageAttributes(wd string, projectId int64, packageName string, b []byte) error {
	packagePath := GetPackagePath(wd, projectId, packageName)
	dotFilePath := path.Join(packagePath, ".steam")

	if err := ioutil.WriteFile(dotFilePath, b, FilePerm); err != nil {
		return fmt.Errorf("Failed writing attributes: %s", err)
	}
	return nil
}

func GetModelPath(wd string, modelId int64) string {
	location := strconv.FormatInt(modelId, 10)
	return path.Join(wd, ModelDir, location)
}

func GetJavaModelPath(wd string, modelId int64, logicalName string) string {
	return path.Join(GetModelPath(wd, modelId), logicalName) + ".java"
}

func GetMOJOPath(wd string, modelId int64, logcialName string) string {
	return path.Join(GetModelPath(wd, modelId), logcialName) + ".zip"
}

func GetWarFilePath(wd string, modelId int64, logicalName string) string {
	return path.Join(GetModelPath(wd, modelId), logicalName) + ".war"
}

func GetPythonWarFilePath(wd string, modelId int64, logicalName string) string {
	return path.Join(GetModelPath(wd, modelId), logicalName) + "_py.war"
}

func GetModelJarFilePath(wd string, modelId int64, logicalName string) string {
	return path.Join(GetModelPath(wd, modelId), logicalName) + ".jar"
}

func GetGenModelPath(wd string, modelId int64) string {
	return path.Join(GetModelPath(wd, modelId), "h2o-genmodel.jar")
}

func GetDeepwaterDepPath(wd string, modelId int64) string {
	return path.Join(GetModelPath(wd, modelId), "deepwater-all.jar")
}

func GetAssetsPath(wd, asset string) string {
	return path.Join(wd, AssetsDir, asset)
}

func GetModelDirs(wd string, modelId int64) ([]string, error) {
	modelDir := GetModelPath(wd, modelId)
	files, err := ioutil.ReadDir(modelDir)
	if err != nil {
		return nil, fmt.Errorf("Model directory read failed: %s: %v", modelDir, err)
	}

	dirs := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}

	return dirs, nil
}

func GetOutPath(wd, jobID string) string {
	return path.Join(wd, OutDir, jobID)
}

func GetTmpFilePath(wd, filename string) string {
	return path.Join(wd, TmpDir, filename)
}

func GetJobLogFilePath(wd, id, suffix string) string {
	return path.Join(wd, LogDir, id+"."+suffix+".log")
}

func GetLogFilePath(wd, name string) string {
	return path.Join(wd, LogDir, name)
}

type Log struct {
	Name       string
	ModifiedAt time.Time
}

type ByModTime []os.FileInfo

func (s ByModTime) Len() int           { return len(s) }
func (s ByModTime) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByModTime) Less(i, j int) bool { return s[i].ModTime().After(s[j].ModTime()) }

func ListLogs(wd string, offset, limit int) ([]*Log, error) {
	logDir := path.Join(wd, LogDir)
	files, err := ioutil.ReadDir(logDir)
	if err != nil {
		return nil, fmt.Errorf("Log directory read failed: %s: %v", logDir, err)
	}

	sort.Sort(ByModTime(files))

	if offset < 0 {
		offset = 0
	}

	if limit < 0 || limit > len(files) {
		limit = len(files)
	}

	logs := make([]*Log, 0)
	for i := offset; i < limit; i++ {
		file := files[i]
		if !file.IsDir() { // should never happen
			logs = append(logs, &Log{file.Name(), file.ModTime()})
		}

	}
	return logs, nil
}

func GetPackPath(wd, kind, pack string) string {
	return path.Join(wd, LibDir, kind, pack+PackExt)
}

func GetBasenameWithoutExt(p string) string {
	basename := path.Base(p)
	return basename[0 : len(basename)-len(path.Ext(basename))]
}

func GetPackDir(packPath string) string {
	return path.Join(path.Dir(packPath), GetBasenameWithoutExt(packPath))
}

func GetPackUrl(host, kind, pack string) string {
	p := fmt.Sprintf("/%s/%s/%s.steam", LibDir, kind, pack)
	return (&url.URL{Scheme: "http", Host: host, Path: p}).String()
}

func GetIP(addr string) string {
	ts := strings.Split(addr, ":")
	if len(ts) > 1 {
		return ts[0]
	}
	return ""
}

func GetExternalIP(fallback string) string {
	ip, err := GetExternalHost()
	if err != nil {
		return fallback
	}
	elems := strings.Split(fallback, ":")
	if len(elems) != 2 {
		return fallback
	}
	return ip + ":" + elems[1]
}

func GetExternalHost() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", fmt.Errorf("Failed determining external IP address.")
}

func Get(url, token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if token != "" {
		req.Header.Set("Authorization", "Basic "+token)
	}

	return http.DefaultClient.Do(req)
}

func Download(token, p, u string, preserveFilename bool) (int64, string, error) {
	res, err := Get(u, token)
	if err != nil {
		return 0, "", fmt.Errorf("File download failed: %s: %v", u, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return 0, "", fmt.Errorf("File download failed: %s: %v:", u, err)
		}
		return 0, "", fmt.Errorf("File download failed: %s: %s / %s", u, res.Status, string(body))
	}

	if preserveFilename {
		disposition := res.Header.Get("Content-Disposition")
		if disposition == "" {
			return 0, "", fmt.Errorf("File download failed: missing Content-Disposition response header")
		}
		_, params, err := mime.ParseMediaType(disposition)
		if err != nil {
			return 0, "", fmt.Errorf("File download failed: could not parse media type from Content-Disposition header: %s : %v", u, err)
		}
		filename, ok := params["filename"]
		if !ok {
			return 0, "", fmt.Errorf("File download failed: missing filename in Content-Disposition header: %s", u)
		}

		p = path.Join(p, filename)
	}

	if err := os.MkdirAll(path.Dir(p), DirPerm); err != nil {
		return 0, "", fmt.Errorf("Download directory creation failed: %s: %v", p, err)
	}

	dst, err := os.Create(p)
	if err != nil {
		return 0, "", fmt.Errorf("Download file creation failed: %s: %v", p, err)
	}
	defer dst.Close()

	if err := os.Chmod(p, FilePerm); err != nil {
		return 0, "", errors.Wrap(err, "setting file permissions")
	}

	size, err := io.Copy(dst, res.Body)
	if err != nil {
		return 0, "", fmt.Errorf("Download file copy failed: %s to %s: %v", u, p, err)
	}

	return size, p, nil
}

func Tail(filename string, lines int) (string, error) {
	if lines <= 0 {
		lines = 100
	} else if lines > 1000 {
		lines = 1000
	}
	b, err := exec.Command("/usr/bin/tail", "-"+strconv.Itoa(lines), filename).Output()
	if err != nil {
		return "", fmt.Errorf("Log tail failed: %v", err)
	}
	return string(b), nil
}
