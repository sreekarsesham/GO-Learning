// this progrms reads a csv to dataframe and filters it for SVT 
// converts that to slice

package main;
import (

	"fmt"
	"os"
	"log"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"runtime"
    
    
	
)

func main(){

        PrintMemUsage()

	csvfile, err := os.Open("df_rhythmInfo_supra_123725_1637316837678082.csv")
         if err != nil {
         	log.Fatal(err)
    	 }
 	// Read CSV 
	df := dataframe.ReadCSV(csvfile) 
	fmt.Println("df =>",df) 
	PrintMemUsage()
	
	// filter only beatChar == s
	
	df_rhythmInfo_supra_filterset := df.Filter(
		dataframe.F{
			Colname:    "beatChar",
			Comparator: "==",
			Comparando: "S",
		},
		
	)
	
	fmt.Println("df_rhythmInfo_supra_filterset",df_rhythmInfo_supra_filterset)
	
	// convert one column to list 
	
	df_supra_list :=  df_rhythmInfo_supra_filterset.Col("HRbpm")
	//fmt.Println("df_supra_list HR+> ",df_supra_list)
	
	var nums []int;
	
	
		el := df_supra_list.Elem(0)
		fmt.Printf("el",el)
		d:= el.Float()
		fmt.Println("float of float =>",d)
		//fmt.Println("", nums)
	
	
	for i, item := range nums {
 
		//fmt.Println(item)
		if(item == 3821){
		  fmt.Println("found at index",i)
		}
	}
	
	 sums := []string{"sss","abc","bbc"}
	
	dfNew := dataframe.New(
	
	 series.New(nums, series.Int, "RR Sec"),
	 series.New(sums, series.String, "samples"),
	
        )
        //fmt.Println("dfNew",dfNew)
        
        fmt.Println("selected",dfNew.Elem(1,1))
        
        f, err := os.Create("output.csv")
	if err != nil {
   	 log.Fatal(err)
	}

	dfNew.WriteCSV(f)
	
	 
	
}




func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}
