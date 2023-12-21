package classpath

import (
	"os"
	"strings"
)

type Classpath struct {
  // 通过-Xjre解析自带类
  bootClassPath Entry
  extClassPath Entry
  // 通过-cp或-classpath解析用户class路径
  userClassPath Entry
}

func Parse(jreOption string, cpOption string) *Classpath {
  cp := &Classpath{}
  return cp
}

// 实现方法
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
  return nil, nil, nil
}

// 实现方法
func (self *Classpath) ToString() string {
  str := []string{}
  str = append(str, self.bootClassPath.ToString())
  str = append(str, self.bootClassPath.ToString())
  str = append(str, self.bootClassPath.ToString())
  return strings.Join(str, ";");
}

// 解析jre路径
func (self *Classpath) parseBootAndExtClassPath(jreOption string) {
  // jrePath := parseJreOption(jreOption)

}

// 从参数获取、从当前目录下获取、从环境变量里获取
func parseJreOption(jreOption string) string {
  if jreOption != "" && exists(jreOption) {
    return jreOption
  }
  if exists("./jre") {
    return "./jre"
  }
  if home := os.Getenv("JAVA_HOME"); home != "" {
    return home;
  }
  panic("can not found jre")
}

// 解析用户类路径
func (self *Classpath) parseUserClassPath(cpOption string) {
  if cpOption == "" {
    cpOption = "."
  }
  self.userClassPath = newEntry(cpOption)
}

// 判断文件夹是否存在
func exists(path string) bool {
  _, err := os.Stat(path)
  if err != nil && os.IsNotExist(err) {
    return false
  }
  return true
}
