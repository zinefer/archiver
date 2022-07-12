package archutil

import (
	"os"
	"log"
	"io/ioutil"
	"io/fs"
	"sort"
	"os/exec"
	"path/filepath"
	"golang.org/x/exp/constraints"
	"strings"
	"time"
)

type FileInfoWithChildModTime struct {
	fs.FileInfo
	path string
	childModTime time.Time
}

func (m *FileInfoWithChildModTime) ChildModTime() time.Time {
    if (!m.childModTime.IsZero()) {
		return m.childModTime
	}

	var selfModTime time.Time
		
	err := filepath.Walk(m.path + "/" + m.Name(),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if (selfModTime.IsZero()) {
				selfModTime = info.ModTime()
				return nil
			}

			if (m.childModTime.Unix() < info.ModTime().Unix()) {
				m.childModTime = info.ModTime()
			}

			return nil
		})
	
	if err != nil {
		log.Fatal(err)
	}

	if (m.childModTime.IsZero()) {
		m.childModTime = selfModTime
	}

	return m.childModTime
}


func ListDirectoryByModifiedTimeAsc(path string) []FileInfoWithChildModTime {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var filesWithChildModTime []FileInfoWithChildModTime
	for _, file := range files {
		fileWithChildModTime := FileInfoWithChildModTime{ FileInfo: file, path: path }
		filesWithChildModTime = append(filesWithChildModTime, fileWithChildModTime)
	}

	sort.Slice(filesWithChildModTime, func(i,j int) bool {
		return filesWithChildModTime[i].ChildModTime().Unix() < 
				filesWithChildModTime[j].ChildModTime().Unix()
	})
	
	return filesWithChildModTime
}

func ListDirectory(path string) []fs.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func Move(path string, destination string) {
	err := os.Rename(path, destination)
	if err != nil {
		log.Fatal(err)
	}
}

func PrintDirectory(path string) string {
	cmd := exec.Command("ls", "--color=always", path)
    stdout, err := cmd.CombinedOutput()
	if err != nil {
        log.Fatal(err)
    }
	return strings.Trim(string(stdout), "\n")
}

func max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}