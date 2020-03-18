package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}
func (c *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className += ".class"
	if data, entry, err := c.bootClassPath.ReadClass(className); err == nil {
		return data, entry, nil
	}
	if data, entry, err := c.extClassPath.ReadClass(className); err == nil {
		return data, entry, nil
	}
	return c.userClassPath.ReadClass(className)
}

func (c *ClassPath) String() string {
	return c.userClassPath.String()
}

func (c *ClassPath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJrePath(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClassPath = newWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClassPath = newWildcardEntry(jreExtPath)
}
func (c *ClassPath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	c.userClassPath = newEntry(cpOption)
}

func getJrePath(jrePath string) string {
	if jrePath != "" && exists(jrePath) {
		return jrePath
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder")

}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
