package main
import (
    "bufio"
    "fmt"
    "strings"
    "strconv"
    "os"
)

var (
    maxCmdNum int
    initHealth int
    restoreHealthThreshold int
    restoreHealthCap int
        
    inputDamageHeal string
)

func restoreHealth(chp *int) {
    var restoreHealthCap_ *int
    restoreHealthCap_ = &restoreHealthCap
    *chp = *chp + *restoreHealthCap_
}

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    
    fmt.Scanln(&maxCmdNum, &initHealth, &restoreHealthThreshold, &restoreHealthCap)
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        inputDamageHeal = scanner.Text()
    }
    inputDamageHealList := strings.Split(inputDamageHeal, " ")
    currentHealth := initHealth
    for _, v := range inputDamageHealList {
        damageAndHeal,_ := strconv.Atoi(v)
        currentHealth = currentHealth + damageAndHeal
        if currentHealth <= restoreHealthThreshold {
            restoreHealth(&currentHealth)
        }
    }
    fmt.Println(currentHealth)
}
