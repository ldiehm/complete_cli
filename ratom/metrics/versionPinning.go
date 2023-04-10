package metrics


import (
	"fmt"
	"os"
	"regexp"
	"strings"
)


func VersionPinning(gitDir string) float32 {

        fmt.Println("versionPinning called")
    
        data, err := os.ReadFile(gitDir + "/package.json")
        if err != nil {
            fmt.Println("Didn't find package.json file")
            return 0
        }

        regex_getAllDependencies, _ := regexp.Compile("\"dependencies\": {([^}]*)")
        regex_getIndividualDependencies, _ := regexp.Compile("(\".*\"): \"(.*)\",")
        regex_getVersion, _ := regexp.Compile("(^\\^[1-9])|(^[0-9]$)|(-)|(^[0-9].x$)|(^~[0-9]$)|(^1.0.0)")

        total := 0
        count_bad_dependency := 0
        res := strings.Split(regex_getAllDependencies.FindStringSubmatch(string(data))[1], "\n")

        for j := 0; j < len(res); j++ {

            curr := regex_getIndividualDependencies.FindStringSubmatch(res[j])

            if len(curr) == 3{
                total += 1
                if regex_getVersion.MatchString(curr[2]) {
                    count_bad_dependency += 1
                }
            }
            
        } 
        

        score := float32(count_bad_dependency) / float32(total)
        
        return score
}