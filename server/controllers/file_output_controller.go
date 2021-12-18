package controllers

import (
	"net/http"
)

func GetCsvFile(w http.ResponseWriter, r *http.Request) {
	// records := [][]string{
	// 	{"first_name", "last_name", "username"},
	// 	{"Rob", "Pike", "rob"},
	// 	{"Ken", "Thompson", "ken"},
	// 	{"Robert", "Griesemer", "gri"},
	// }

	// file, err := os.Create("sample.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// wc := csv.NewWriter(w)

	// // w := csv.NewWriter(os.Stdout)

	// for _, record := range records {
	// 	if err := wc.Write(record); err != nil {
	// 		log.Fatalln("error writing record to csv:", err)
	// 	}
	// }

	// wc.Flush()
	// if err := wc.Error(); err != nil {
	// 	log.Fatal(err)
	// }
	// file.Close()

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w.Header().Set("Content-Disposition", "attachment; filename=test.csv")
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Length", string(len(records)))
	var out []byte
	for _, record := range records {
		for _, value := range record {
		  if value == record[len(record) - 1] {
			out = []byte(value)
		  } else {
			out = []byte(value + ",")
		  }
		  w.Write(out)
		}
		w.Write([]byte("\n"))
	}
	// for i := 0; i < 4; i++ {
	// 	out = []byte("test,test2,test3")
	// 	w.Write(out)
	// }
}
