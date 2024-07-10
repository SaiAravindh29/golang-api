package main

import (
	as "Assign2/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Student is a struct
type Student struct {
	Name     string      `json:"name"`
	Marks    as.Marks    `json:"marks"`
	School   as.School   `json:"school"`
	Address  as.Address  `json:"address"`
	Personal as.Personal `json:"personal"`
}

type ResultData interface{}

type ErrorStruct struct {
	Status string
	Error  string
	Result ResultData
}

var Gresponse ErrorStruct

var Students []Student /* this slice is used as main database for the api*/

func main() {
	log.Println("Starting...")
	http.HandleFunc("/getStudent", getStudent)
	http.HandleFunc("/createStudent", createStudent)
	http.HandleFunc("/updateStudent", updateStudent)
	http.ListenAndServe(":10000", nil)
}

/*
Purpose : to get the student details
Request :

   Body : nil
----------------------
   Header : nil
----------------------
   Params :
	  name : "Sai"
	 dtype : "Marks"
----------------------
	Response :
	On Success
	==========
	{
		status : s
		 error : ""
		result : "message"
	}
	On Error
	=========
	{
	   status : e
		error : error msg
	   result : nil
	}

Authorization : SAI ARAVINDH
Date : 08-07-2024 */

func getStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting data ...")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Content-Type", "application/json")

	if r.Method == "GET" { // to ensure that this func should work only the method is fixed in GET
		Gresponse.Status = "s"
		name := r.URL.Query().Get("name")
		dtype := r.URL.Query().Get("dtype")

		// if the database is empty it will throw a message
		if Students == nil {
			Gresponse.Error = "No Data Found : Empty Database "
			Gresponse.Status = "e"
			lresponse, _ := json.Marshal(Gresponse)
			fmt.Fprintln(w, string(lresponse))
			return
		}

		if name == "" && dtype != "" { // used to filter the data based on the type and not considering the name

			if dtype == "All" || dtype == "all" {
				Gresponse.Status = "s"
				Gresponse.Error = ""
				Gresponse.Result = Students
				lresponse, err := json.Marshal(Gresponse)
				if err != nil {
					Gresponse.Status = "e"
					Gresponse.Error = "AGS01 " + err.Error()
					lresponse, _ := json.Marshal(Gresponse)
					fmt.Fprintln(w, string(lresponse))
					return
				}
				fmt.Fprintln(w, string(lresponse))
				Gresponse.Result = nil

			} else if dtype == "Name" || dtype == "name" { // to get the Names of all the students

				var NameData []string
				for _, student := range Students {
					NameData = append(NameData, student.Name)
				}
				Gresponse.Status = "s"
				Gresponse.Error = ""
				Gresponse.Result = NameData
				lresponse, err := json.Marshal(Gresponse)
				if err != nil {
					Gresponse.Status = "e"
					Gresponse.Error = "AGS02 " + err.Error()
					lresponse, _ := json.Marshal(Gresponse)
					fmt.Fprintln(w, string(lresponse))
					return
				}
				fmt.Fprintln(w, string(lresponse))
				Gresponse.Result = nil

			} else if dtype == "Marks" || dtype == "marks" { // to get marks of all the students

				type MarksData struct {
					Name  string   `json:"name"`
					Marks as.Marks `json:"marks"`
				}

				var markData MarksData
				var AllMarks []MarksData

				for i := 0; i < len(Students); i++ {

					markData.Name = Students[i].Name
					markData.Marks = Students[i].Marks
					AllMarks = append(AllMarks, markData)
				}

				Gresponse.Status = "s"
				Gresponse.Error = ""
				Gresponse.Result = AllMarks
				lresponse, err := json.Marshal(Gresponse)
				if err != nil {
					Gresponse.Status = "e"
					Gresponse.Error = "AGS03 " + err.Error()
					lresponse, _ := json.Marshal(Gresponse)
					fmt.Fprintln(w, string(lresponse))
					return
				}
				fmt.Fprintln(w, string(lresponse))
				Gresponse.Result = nil

			} else if dtype == "School" || dtype == "school" { // to get School Details of all the students

				type SchoolData struct {
					Name   string    `json:"name"`
					School as.School `json:"school"`
				}

				var schoolData SchoolData
				var AllSchool []SchoolData
				for i := 0; i < len(Students); i++ {
					schoolData.Name = Students[i].Name
					schoolData.School = Students[i].School
					AllSchool = append(AllSchool, schoolData)
				}

				Gresponse.Status = "s"
				Gresponse.Error = ""
				Gresponse.Result = AllSchool
				lresponse, err := json.Marshal(Gresponse)
				if err != nil {
					Gresponse.Status = "e"
					Gresponse.Error = "AGS04 " + err.Error()
					lresponse, _ := json.Marshal(Gresponse)
					fmt.Fprintln(w, string(lresponse))
					return
				}
				fmt.Fprintln(w, string(lresponse))
				Gresponse.Result = nil

			} else if dtype == "Address" || dtype == "address" { // to get Address Details of all the students

				type AddressData struct {
					Name    string     `json:"name"`
					Address as.Address `json:"address"`
				}

				var addressData AddressData
				var AllAddress []AddressData
				for i := 0; i < len(Students); i++ {

					addressData.Name = Students[i].Name
					addressData.Address = Students[i].Address
					AllAddress = append(AllAddress, addressData)

				}
				Gresponse.Status = "s"
				Gresponse.Error = ""
				Gresponse.Result = AllAddress
				lresponse, err := json.Marshal(Gresponse)
				if err != nil {
					Gresponse.Status = "e"
					Gresponse.Error = "AGS05 " + err.Error()
					lresponse, _ := json.Marshal(Gresponse)
					fmt.Fprintln(w, string(lresponse))
					return
				}
				fmt.Fprintln(w, string(lresponse))
				Gresponse.Result = nil

			} else if dtype == "Personal" || dtype == "personal" { // to get Personal Details of all the students

				type PersonalData struct {
					Name     string      `json:"name"`
					Personal as.Personal `json:"personal"`
				}

				var personalData PersonalData
				var AllPersonal []PersonalData
				for i := 0; i < len(Students); i++ {

					personalData.Name = Students[i].Name
					personalData.Personal = Students[i].Personal
					AllPersonal = append(AllPersonal, personalData)

				}
				Gresponse.Status = "s"
				Gresponse.Error = ""
				Gresponse.Result = AllPersonal
				lresponse, err := json.Marshal(Gresponse)
				if err != nil {
					Gresponse.Status = "e"
					Gresponse.Error = "AGS06 " + err.Error()
					lresponse, _ := json.Marshal(Gresponse)
					fmt.Fprintln(w, string(lresponse))
					return
				}

				fmt.Fprintln(w, string(lresponse))
				Gresponse.Result = nil

			} else {
				Gresponse.Status = "e"
				Gresponse.Error = "Invalid data type"
				lresponse, _ := json.Marshal(Gresponse)
				fmt.Fprintln(w, string(lresponse))
				// http.Error(w, "Invalid data type", http.StatusBadRequest)
				return
			}

			/**************************************************************************************************************/
		} else if name != "" && dtype != "" { // to get Details of specific student

			var NameVerify bool = false

			// this range will check the database , whether the entered name is present or not, if not it will throw a error msg
			for _, s := range Students {
				if s.Name == name {
					NameVerify = true
				}
			}

			if NameVerify { // if name is presented, this block is executed.

				for i := range Students {

					if Students[i].Name == name {
						if dtype == "All" || dtype == "all" {
							// var tempDetails Student
							tempDetails := Students[i] // using the temporary variable to to get the list of all the details of the selected student.
							Gresponse.Status = "s"
							Gresponse.Error = ""
							Gresponse.Result = tempDetails
							lresponse, err := json.Marshal(Gresponse)
							if err != nil {
								Gresponse.Status = "e"
								Gresponse.Error = "AGS07 " + err.Error()
								lresponse, _ := json.Marshal(Gresponse)
								fmt.Fprintln(w, string(lresponse))
								return
							}
							fmt.Fprintln(w, string(lresponse))
							Gresponse.Result = nil
							return
						} else if dtype == "Marks" || dtype == "marks" {

							type MarksData struct {
								Name  string   `json:"name"`
								Marks as.Marks `json:"marks"`
							}

							var markData MarksData
							markData.Name = Students[i].Name
							markData.Marks = Students[i].Marks

							Gresponse.Status = "s"
							Gresponse.Error = ""
							Gresponse.Result = markData
							lresponse, err := json.Marshal(Gresponse)
							if err != nil {
								Gresponse.Status = "e"
								Gresponse.Error = "AGS08 " + err.Error()
								lresponse, _ := json.Marshal(Gresponse)
								fmt.Fprintln(w, string(lresponse))
								return
							}
							fmt.Fprintln(w, string(lresponse))
							Gresponse.Result = nil
							return
						} else if dtype == "School" || dtype == "school" {

							type SchoolData struct {
								Name   string    `json:"name"`
								School as.School `json:"school"`
							}

							var schoolData SchoolData
							schoolData.Name = Students[i].Name
							schoolData.School = Students[i].School

							Gresponse.Status = "s"
							Gresponse.Error = ""
							Gresponse.Result = schoolData
							lresponse, err := json.Marshal(Gresponse)
							if err != nil {
								Gresponse.Status = "e"
								Gresponse.Error = "AGS09 " + err.Error()
								lresponse, _ := json.Marshal(Gresponse)
								fmt.Fprintln(w, string(lresponse))
								return
							}
							fmt.Fprintln(w, string(lresponse))
							Gresponse.Result = nil
							return
						} else if dtype == "Address" || dtype == "address" {

							type AddressData struct {
								Name    string     `json:"name"`
								Address as.Address `json:"address"`
							}

							var addressData AddressData
							addressData.Name = Students[i].Name
							addressData.Address = Students[i].Address

							Gresponse.Status = "s"
							Gresponse.Error = ""
							Gresponse.Result = addressData
							lresponse, err := json.Marshal(Gresponse)
							if err != nil {
								Gresponse.Status = "e"
								Gresponse.Error = "AGS10 " + err.Error()
								lresponse, _ := json.Marshal(Gresponse)
								fmt.Fprintln(w, string(lresponse))
								return
							}
							fmt.Fprintln(w, string(lresponse))
							Gresponse.Result = nil
							return
						} else if dtype == "Personal" || dtype == "personal" {
							type PersonalData struct {
								Name     string      `json:"name"`
								Personal as.Personal `json:"personal"`
							}

							var personalData PersonalData
							personalData.Name = Students[i].Name
							personalData.Personal = Students[i].Personal

							Gresponse.Status = "s"
							Gresponse.Error = ""
							Gresponse.Result = personalData
							lresponse, err := json.Marshal(Gresponse)
							if err != nil {
								Gresponse.Status = "e"
								Gresponse.Error = "AGS11 " + err.Error()
								lresponse, _ := json.Marshal(Gresponse)
								fmt.Fprintln(w, string(lresponse))
								return
							}
							fmt.Fprintln(w, string(lresponse))
							Gresponse.Result = nil
							return
						} else {
							Gresponse.Status = "e"
							Gresponse.Error = "Invalid data type"
							lresponse, _ := json.Marshal(Gresponse)
							fmt.Fprintln(w, string(lresponse))
							return
						}
					}
				}

			} else {
				Gresponse.Status = "e"
				Gresponse.Error = "StudentName is not Present in the database"
				lresponse, _ := json.Marshal(Gresponse)
				fmt.Fprintln(w, string(lresponse))
				NameVerify = true
				return

			}

		} else {
			Gresponse.Status = "e"
			Gresponse.Error = "Invalid data type"
			lresponse, _ := json.Marshal(Gresponse)
			fmt.Fprintln(w, string(lresponse))
			return
		}

	} else {
		Gresponse.Status = "e"
		Gresponse.Error = "Invalid Method"
		lresponse, _ := json.Marshal(Gresponse)
		fmt.Fprintln(w, string(lresponse))
		return
	}

}

/*
Purpose : to Create the student details
Request :
   Body :
    {
   		marks    :
		school   :
		address  :
		personal :
    }
----------------------
   Header :
   {
     name :
   }
----------------------
   Params :
----------------------
Response :
	On Success
	==========
	{
		status : s
		 error : ""
		result : "message"
	}
	On Error
	=========
	{
	   status : e
		error : error msg
	   result : nil
	}
Authorization : SAI ARAVINDH
Date : 08-07-2024 */

func createStudent(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Creating Student .........")
	// fmt.Fprintln(w, len(Students))
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {

		namekey := r.Header.Get("name")

		Studentcheck := 0
		defer r.Body.Close()
		var newStudent Student
		newStudent.Name = namekey

		err := json.NewDecoder(r.Body).Decode(&newStudent)
		if err != nil {
			log.Println("Error parsing request body:", err)
			Gresponse.Status = "e"
			Gresponse.Error = "ACS01 " + err.Error()
			lresponse, _ := json.Marshal(Gresponse)
			fmt.Fprintln(w, string(lresponse))
			return
		}

		if len(Students) == 0 {

			if newStudent.Name != "" {

				if newStudent.Marks.M10.English != 0 && newStudent.Marks.M10.Tamil != 0 && newStudent.Marks.M10.Maths != 0 && newStudent.Marks.M10.Science != 0 && newStudent.Marks.M10.Social != 0 {

					if newStudent.School.S10.SchoolName != "" && newStudent.School.S10.Address != "" && newStudent.School.S10.Place != "" && newStudent.School.S10.Pincode != "" && newStudent.School.S10.Type != "" {

						if newStudent.Address.Landmark != "" && newStudent.Address.Pincode != "" && newStudent.Address.StudentAddress != "" {

							if newStudent.Personal.FatherName != "" && newStudent.Personal.MotherName != "" && newStudent.Personal.Age != 0 && newStudent.Personal.Gender != "" {

								Students = append(Students, newStudent)
								Gresponse.Status = "s"
								Gresponse.Error = ""
								Gresponse.Result = "Student " + newStudent.Name + " is Created"
								lresponse, err := json.Marshal(Gresponse)
								if err != nil {
									Gresponse.Status = "e"
									Gresponse.Error = "ACS01 " + err.Error()
									lresponse, _ := json.Marshal(Gresponse)
									fmt.Fprintln(w, string(lresponse))
									return
								}
								fmt.Fprintln(w, string(lresponse))
								Gresponse.Result = nil
								return
							}
						}
					}

				}

			}
		} else if len(Students) != 0 {

			for index := 0; index < len(Students); index++ {
				if Students[index].Name == namekey {
					Studentcheck++

				}

			}

			if Studentcheck == 0 {

				if newStudent.Name != "" {

					if newStudent.Marks.M10.English != 0 && newStudent.Marks.M10.Tamil != 0 && newStudent.Marks.M10.Maths != 0 && newStudent.Marks.M10.Science != 0 && newStudent.Marks.M10.Social != 0 {

						if newStudent.School.S10.SchoolName != "" && newStudent.School.S10.Address != "" && newStudent.School.S10.Place != "" && newStudent.School.S10.Pincode != "" && newStudent.School.S10.Type != "" {

							if newStudent.Address.Landmark != "" && newStudent.Address.Pincode != "" && newStudent.Address.StudentAddress != "" {

								if newStudent.Personal.FatherName != "" && newStudent.Personal.MotherName != "" && newStudent.Personal.Age != 0 && newStudent.Personal.Gender != "" {

									Students = append(Students, newStudent)
									Gresponse.Status = "s"
									Gresponse.Error = ""
									Gresponse.Result = "Student " + newStudent.Name + " is Created"
									lresponse, _ := json.Marshal(Gresponse)
									if err != nil {
										Gresponse.Status = "e"
										Gresponse.Error = "ACS01 " + err.Error()
										lresponse, _ := json.Marshal(Gresponse)
										fmt.Fprintln(w, string(lresponse))
										return
									}
									fmt.Fprintln(w, string(lresponse))
									Gresponse.Result = nil
									// fmt.Fprintf(w, "\nStudent %v created", newStudent.Name)
									Studentcheck = 0
									// fmt.Fprintln(w, Students)
									return
								}
							}
						}

					}

				}
			} else {
				Gresponse.Status = "e"
				Gresponse.Error = "Student " + namekey + " is already Presented ...:)"
				lresponse, _ := json.Marshal(Gresponse)
				fmt.Fprintln(w, string(lresponse))
				return
			}

		}
		Gresponse.Status = "e"
		Gresponse.Error = "Student Not created"
		lresponse, _ := json.Marshal(Gresponse)
		fmt.Fprintln(w, string(lresponse))
		Studentcheck = 0
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	log.Println("Error reading request body:", err)
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// err = json.Unmarshal(body, &newStudent)
		// if err != nil {
		// 	log.Println("Error parsing request body:", err)
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// Students = append(Students, newStudent)
	} else {
		Gresponse.Status = "e"
		Gresponse.Error = "Invalid Method"
		lresponse, _ := json.Marshal(Gresponse)
		fmt.Fprintln(w, string(lresponse))
		return
	}
}

/*
Purpose : to Update the student details
Request :
   Body :
    {
   		marks    :
		school   :
		address  :
		personal :
    }
----------------------
   Header :
   {
     name :
   }
----------------------
   Params :
	  name : "Sai"
	 dtype : "Marks"
----------------------
Response :
	On Success
	==========
	{
		status : s
		 error : ""
		result : "message"
	}
	On Error
	=========
	{
	   status : e
		error : error msg
	   result : nil
	}
Authorization : SAI ARAVINDH
Date : 08-07-2024 */

func updateStudent(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Updating Student .........")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {

		defer r.Body.Close()
		studentCheck := 0
		name := r.URL.Query().Get("name")
		dtype := r.URL.Query().Get("dtype")

		for i := 0; i < len(Students); i++ {

			if Students[i].Name == name {

				if dtype == "Marks" || dtype == "marks" {

					var tempMarks as.Marks

					body, err := ioutil.ReadAll(r.Body)

					if err != nil {
						log.Println("Error reading request body:", err)
						Gresponse.Status = "e"
						Gresponse.Error = "AUS01 " + err.Error()
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

					err = json.Unmarshal(body, &tempMarks)
					if err != nil {
						log.Println("Error parsing request body:", err)
						Gresponse.Status = "e"
						Gresponse.Error = "AUS01 " + err.Error()
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

					if tempMarks.M10 == nil {
						Gresponse.Status = "e"
						Gresponse.Error = "Invalid Details entered."
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

					if tempMarks.M10.English > 0 && tempMarks.M10.Tamil > 0 && tempMarks.M10.Maths > 0 && tempMarks.M10.Science > 0 && tempMarks.M10.Social > 0 {

						Students[i].Marks = tempMarks
						message := Students[i].Name + " Marks is Updated Successfully"
						studentCheck = 0
						// jsonData, err := json.Marshal(message)
						Gresponse.Status = "s"
						Gresponse.Error = ""
						Gresponse.Result = message
						lresponse, err := json.Marshal(Gresponse)
						if err != nil {
							Gresponse.Status = "e"
							Gresponse.Error = "AUS01 " + err.Error()
							lresponse, _ := json.Marshal(Gresponse)
							fmt.Fprintln(w, string(lresponse))
							return
						}
						fmt.Fprintln(w, string(lresponse))
						Gresponse.Result = nil
						return
					} else {
						Gresponse.Status = "e"
						Gresponse.Error = "Student Name Already Presented"
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

				} else if dtype == "name" || dtype == "Name" {

					oldname := Students[i].Name

					tempname := r.Header.Get("name")

					if tempname == "" {
						Gresponse.Status = "e"
						Gresponse.Error = "Data not entered in header ."
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}
					var updateNamesCheck bool = true

					for names := range Students {
						if tempname == Students[names].Name {
							updateNamesCheck = false
						}
					}

					if updateNamesCheck {
						Students[i].Name = tempname

						message := oldname + " changed to " + Students[i].Name + " **  Updated Successfully **"
						studentCheck = 0
						jsonData, err := json.Marshal(message)
						if err != nil {
							Gresponse.Status = "e"
							Gresponse.Error = "AUS02 " + err.Error()
							lresponse, _ := json.Marshal(Gresponse)
							fmt.Fprintln(w, string(lresponse))
							return
						}
						Gresponse.Status = "s"
						Gresponse.Error = ""
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						fmt.Fprintln(w, string(jsonData))
						return
					} else {
						Gresponse.Status = "e"
						Gresponse.Error = "Student Name Already Presented"
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

				} else if dtype == "School" || dtype == "school" {

					var tempSchool as.School

					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						log.Println("Error reading request body:", err)
						Gresponse.Status = "e"
						Gresponse.Error = "AUS03 " + err.Error()
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

					err1 := json.Unmarshal(body, &tempSchool)
					fmt.Printf("Type of school is  %T", tempSchool)
					if err1 != nil {
						log.Println("Error parsing request body:", err1)
						Gresponse.Status = "e"
						Gresponse.Error = "AUS03 " + err1.Error()
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

					if tempSchool.S10 == nil {
						Gresponse.Status = "e"
						Gresponse.Error = "Invalid Details entered."
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

					if tempSchool.S10.SchoolName != "" && tempSchool.S10.Place != "" &&
						tempSchool.S10.Address != "" && tempSchool.S10.Pincode != "" && tempSchool.S10.Type != "" {
						Students[i].School = tempSchool
						message := Students[i].Name + " School Details is Updated Successfully"
						studentCheck = 0

						Gresponse.Status = "s"
						Gresponse.Error = ""
						Gresponse.Result = message
						lresponse, err := json.Marshal(Gresponse)
						if err != nil {
							Gresponse.Status = "e"
							Gresponse.Error = "AUS03 " + err.Error()
							lresponse, _ := json.Marshal(Gresponse)
							fmt.Fprintln(w, string(lresponse))
							return
						}
						fmt.Fprintln(w, string(lresponse))
						Gresponse.Result = nil
						return
					} else {
						Gresponse.Status = "e"
						Gresponse.Error = "Some Mandatory detais are not entered."
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

				} else if dtype == "Address" || dtype == "address" {

					var tempAddress as.Address
					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						log.Println("Error reading request body:", err)
						Gresponse.Status = "e"
						Gresponse.Error = "AUS04 " + err.Error()
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}
					err = json.Unmarshal(body, &tempAddress)
					if err != nil {
						log.Println("Error parsing request body:", err)
						Gresponse.Status = "e"
						Gresponse.Error = "AUS04 " + err.Error()
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

					if tempAddress.StudentAddress != "" && tempAddress.Landmark != "" && tempAddress.Pincode != "" {
						Students[i].Address = tempAddress
						message := Students[i].Name + " Address is Updated Successfully"
						studentCheck = 0

						Gresponse.Status = "s"
						Gresponse.Error = ""
						Gresponse.Result = message
						lresponse, err := json.Marshal(Gresponse)
						if err != nil {
							Gresponse.Status = "e"
							Gresponse.Error = "AUS04 " + err.Error()
							lresponse, _ := json.Marshal(Gresponse)
							fmt.Fprintln(w, string(lresponse))
							return
						}
						fmt.Fprintln(w, string(lresponse))
						Gresponse.Result = nil
						return
					} else {
						Gresponse.Status = "e"
						Gresponse.Error = "Some Mandatory detais are not entered."
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

				} else if dtype == "Personal" || dtype == "personal" {

					var tempPersonal as.Personal
					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						log.Println("Error reading request body:", err)
						Gresponse.Status = "e"
						Gresponse.Error = "AUS05 " + err.Error()
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}
					err = json.Unmarshal(body, &tempPersonal)
					if err != nil {
						log.Println("Error parsing request body:", err)
						Gresponse.Status = "e"
						Gresponse.Error = "AUS05 " + err.Error()
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

					if tempPersonal.FatherName != "" && tempPersonal.MotherName != "" && tempPersonal.Age > 0 && tempPersonal.Gender != "" {
						Students[i].Personal = tempPersonal
						message := Students[i].Name + " Personal Details is Updated Successfully"
						studentCheck = 0

						Gresponse.Status = "s"
						Gresponse.Error = ""
						Gresponse.Result = message
						lresponse, err := json.Marshal(Gresponse)
						if err != nil {
							Gresponse.Status = "e"
							Gresponse.Error = "AUS05 " + err.Error()
							lresponse, _ := json.Marshal(Gresponse)
							fmt.Fprintln(w, string(lresponse))
							return
						}
						fmt.Fprintln(w, string(lresponse))
						Gresponse.Result = nil
						return
					} else {
						Gresponse.Status = "e"
						Gresponse.Error = "Some Mandatory detais are not entered."
						lresponse, _ := json.Marshal(Gresponse)
						fmt.Fprintln(w, string(lresponse))
						return
					}

				} else {
					Gresponse.Status = "e"
					Gresponse.Error = "Invalid data type"
					lresponse, _ := json.Marshal(Gresponse)
					fmt.Fprintln(w, string(lresponse))
					studentCheck = 0
					return
				}

			} else {
				studentCheck++
			}
		}

		if studentCheck != 0 {

			Gresponse.Status = "e"
			Gresponse.Error = "Invalid data type"
			lresponse, _ := json.Marshal(Gresponse)
			fmt.Fprintln(w, string(lresponse))
			studentCheck = 0
			return
		}

	} else {
		Gresponse.Status = "e"
		Gresponse.Error = "Invalid Method"
		lresponse, _ := json.Marshal(Gresponse)
		fmt.Fprintln(w, string(lresponse))
		return
	}

}
