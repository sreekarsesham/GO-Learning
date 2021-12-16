// this progrms reads a csv to dataframe and filters it for SVT 
// converts that to slice

package main;
import (

	"fmt"
	
	 "github.com/tobgu/qframe"
	
	"runtime"
	"os"
	"log"
    
    
	
)

func main(){

        PrintMemUsage()

	// 1 . read from csv
	csvfile, err := os.Open("df_rhythmInfo_supra_123725_1637316837678082.csv")
	if err != nil {
    		log.Fatal(err)
	}
	f := qframe.ReadCSV(csvfile)
	fmt.Printf("%T",f)
	fmt.Printf("memory taken to create a df")
	 PrintMemUsage()
	 
	 // 2 . filtering
	 
	 newF := f.Filter(
         qframe.Filter{Column: "beatChar", Comparator: "=", Arg: "S"},)
         
         fmt.Println("filter==>",newF)
         PrintMemUsage()
         
         // convert 1 column to list 
         newFSlice := newF.Slice(0,1)
         //fmt.Println(newFSlice.Select("RowNumber"))
         
        // var beatStr []string;
         
         v := newFSlice.MustIntView("RowNumber")
         fmt.Println("v = =",len(v.Slice()))
         var s int;
         for i:=0; i<newFSlice.Len();i++{
           s = v.Slice()[i] 
           fmt.Println("s====",s)
         }
         
         //fmt.Println("v.Slice().len====",v.Slice().Len())
         
         // creating a new df
         
          newCol:= []int{1, 2, 2, 3, 3}
          newCol2:= []int{9, 7, 0, 3, 3}
         
         newDf := qframe.New(map[string]interface{} {
         
         	"nums":newCol,
         })
         fmt.Println("newDf",newDf)
         
         newDf2 := newDf
         
         
         sel:= newDf2.Select("nums")
         sel = newCol2
         fmt.Println("sel",sel)
         
	
	PrintMemUsage()
	
	
	
         
         

	
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
