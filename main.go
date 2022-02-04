package main

import "github.com/tarunKoyalwar/matthew/cmd"

//Now  Webchecklist is not working fix that
//Best way is to use Burpsuite

func main() {
	cmd.Execute()
}

// func main() {
// 	web := flag.Bool("web", false, "Subdomains Checklist page")
// 	sub := flag.String("sub", "", "Update Checklist of this subdomain")
// 	working := flag.Bool("working", false, "Get Name of Working Subdomain under Web Checklist")
// 	org := flag.Bool("org", false, "Organization Checklist page")
// 	tname := flag.String("item", "", "ToolName/CheckList Item (Part of It would do)")
// 	checklist := flag.Bool("checklist", false, "Only Print CheckList (Can only be used with web|org)")
// 	pname := flag.String("pname", "", "Page Name Ex : in-scope etc")
// 	post := flag.Bool("post", false, "Will Post data to destination\n\tIf not used returns data from source")
// 	append := flag.Bool("append", false, "Append data to existing values")
// 	save := flag.Bool("save", false, "Save data to MongoDB")
// 	proxy := flag.String("proxy", "", "Proxy URL")
// 	chunk := flag.Int("chunk", 10, "Chunk Size while sending to network")
// 	// os.Setenv("HTTP_PROXY", "http://127.0.0.1:8080/")

// 	flag.Parse()

// 	if *proxy != "" {
// 		restapi.Proxy = *proxy
// 	}

// 	stdin.ChunkSize = *chunk

// 	var ch chan stdin.Receive

// 	// s := neo.NewStreamWriterWithCondition(4)

// 	if stdin.CheckStdin() {
// 		//has stdin
// 		ch = stdin.GetStdinPipe()
// 		defer stdin.Wg.Wait()

// 	}

// 	//Post Data Not yet created
// 	_ = *append

// 	r, err := restapi.NewRestapi("http://127.0.0.1:8088")
// 	HandleError(err)
// 	r.Alive()
// 	if *working {
// 		dat, err := r.Workingsub()
// 		if dat == "default" {
// 			fmt.Println("Working Subdomain Not Set in UI")
// 			return
// 		} else {
// 			PrintAndHandleError("Working Subdomain is", dat, err)
// 		}
// 	} else if *save {
// 		HandleError(r.SavetoDB())
// 	} else if *pname != "" {
// 		if *post {

// 			func(chx chan stdin.Receive) {
// 				start := true
// 				for {
// 					dat, ok := <-ch
// 					if !ok {
// 						break
// 					} else {
// 						if start {
// 							r.PostPage(*pname, dat.ByteData, *append)
// 							fmt.Print(dat.StringData)
// 							start = false
// 						} else {
// 							r.PostPage(*pname, dat.ByteData, true)
// 							fmt.Print(dat.StringData)
// 						}
// 					}

// 				}
// 				r.SavetoDB()
// 			}(ch)

// 		} else {
// 			dat, err := r.GetPage(*pname)
// 			HandleError(err)
// 			fmt.Println(dat)
// 		}
// 	} else if *web {
// 		if *checklist {
// 			//print checklist
// 			bin, err := r.GetCheckList(true)
// 			HandleError(err)
// 			restapi.PrettyPrint(bin)
// 			return
// 		}
// 		if *sub != "" && *tname != "" {
// 			if *post {
// 				func(chx chan stdin.Receive) {
// 					start := true
// 					for {
// 						dat, ok := <-ch
// 						if !ok {
// 							break
// 						} else {
// 							if start {
// 								r.PostWebToolOutput(*sub, *tname, dat.ByteData, *append)
// 								fmt.Print(dat.StringData)
// 								start = false
// 							} else {
// 								r.PostWebToolOutput(*sub, *tname, dat.ByteData, true)
// 								fmt.Print(dat.StringData)
// 							}

// 						}

// 					}
// 					r.SavetoDB()
// 				}(ch)

// 			} else {
// 				dat, err := r.GetWebToolOutput(*sub, *tname)
// 				HandleError(err)
// 				fmt.Println(dat)
// 			}
// 		} else {
// 			fmt.Println("Subdomain Not Mentioned or Toolname not mentioned")
// 		}
// 	} else if *org {
// 		if *checklist {
// 			//print checklist
// 			bin, err := r.GetCheckList(false)
// 			HandleError(err)
// 			restapi.PrettyPrint(bin)
// 			return
// 		}
// 		if *tname != "" {
// 			if *post {
// 				//post data here
// 				func(chx chan stdin.Receive) {
// 					start := true
// 					for {
// 						dat, ok := <-ch
// 						if !ok {
// 							break
// 						} else {
// 							if start {
// 								r.PostOrgToolOutput(*tname, dat.ByteData, *append)
// 								fmt.Print(dat.StringData)
// 								start = false
// 							} else {
// 								r.PostOrgToolOutput(*tname, dat.ByteData, true)
// 								fmt.Print(dat.StringData)
// 							}

// 						}

// 					}
// 					r.SavetoDB()
// 				}(ch)
// 			} else {
// 				dat, err := r.GetOrgToolOutput(*tname)
// 				HandleError(err)
// 				fmt.Println(dat)
// 			}
// 		} else {
// 			fmt.Println("Toolname not mentioned")
// 		}
// 	} else {
// 		fmt.Println("Looks Like You did not select any option use -h")
// 	}

// 	// pagename := flag.String("page", "", "PageName ex: in-scope etc")
// 	// append := flag.Bool("append", false, "Append to existing data")

// 	// flag.Parse()

// 	// ch := stdin.GetStdinPipe()
// 	// defer stdin.Wg.Wait()

// 	// s := neo.NewStreamWriter()

// 	// if stdin.CheckStdin() {

// 	// }

// 	// // fmt.Println("sending")

// 	// r, er := restapi.NewRestapi("http://127.0.0.1:8088/")
// 	// HandleError(er)

// 	// // fmt.Println("startedapi")

// 	// // data := s.Buffer.Bytes()

// 	// // fmt.Printf("Buffered data is %v\n", string(data))

// 	// dat, er2 := r.PostPage(*pagename, s.Buffer.Bytes(), *append)

// 	// HandleError(er2)
// 	// fmt.Println(dat)

// 	// // r, err := runner.GetcmdStruct("ping -c 10 google.com", "")
// 	// // if err != nil {
// 	// // 	panic(err)
// 	// // }

// 	// // r.Stdout = s

// 	// // er := r.Run()
// 	// // if er != nil {
// 	// // 	panic(er)
// 	// // }

// 	// // fmt.Println("Final Data")
// 	// // fmt.Println(s.Buffer.String())
// }

// func HandleError(er error) {
// 	if er != nil {
// 		panic(er)
// 	}
// }

// func PrintAndHandleError(msg string, dat string, err error) {
// 	HandleError(err)
// 	fmt.Printf("%v :\n%v", msg, dat)
// }
