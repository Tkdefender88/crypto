package main

import "fmt"

var (
	pt = "1815907181716474805136452061793917684000871911998851410864797078911161933431337632774829806207517001958179617856720738101327521552576351369691667910371502971480153619360010341709624631317220940851114914911751279825748"
	n  = "29129463609326322559521123136222078780585451208149138547799121083622333250646678767769126248182207478527881025116332742616201890576280859777513414460842754045651093593251726785499360828237897586278068419875517543013545369871704159718105354690802726645710699029936754265654381929650494383622583174075805797766685192325859982797796060391271817578087472948205626257717479858369754502615173773514087437504532994142632207906501079835037052797306690891600559321673928943158514646572885986881016569647357891598545880304236145548059520898133142087545369179876065657214225826997676844000054327141666320553082128424707948750331"
	e  = 3
)

func main() {
	fmt.Println(len(pt), len(n))
}
