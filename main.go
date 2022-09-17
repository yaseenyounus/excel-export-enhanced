package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lukasjarosch/go-docx"
	"github.com/xuri/excelize/v2"
)

type registrantInfo struct {
	fatherName       string
	motherName       string
	address          string
	phoneNumber      string
	email            string
	emergencyContact string
	emergencyNumber  string
	student1Name     string
	student1DOB      string
	student2Name     string
	student2DOB      string
	student3Name     string
	student3DOB      string
	student4Name     string
	student4DOB      string
}

var cleanRegistrantData []registrantInfo

var excelFileName string
var sheetName string

var starLabels = []string{
	"ParentInformation_FathersName_First",
	"ParentInformation_FathersName_Last",
	"ParentInformation_MothersName_First",
	"ParentInformation_MothersName_Last",
	"ParentInformation_Address_Line1",
	"ParentInformation_Address_City",
	"ParentInformation_Address_State",
	"ParentInformation_Address_PostalCode",
	"ParentInformation_Phone",
	"ParentInformation_Email",
	"EmergencyContactInformation_EmergencyContactName_First",
	"EmergencyContactInformation_EmergencyContactName_Last",
	"EmergencyContactInformation_EmergencyContactPhone",
	"StudentInformation_FirstStudentsInformation_FristStudentsName_First",
	"StudentInformation_FirstStudentsInformation_FristStudentsName_Last",
	"StudentInformation_FirstStudentsInformation_DateOfBirth",
	"StudentInformation_SecondStudentsInformation_SecondStudentsName_First",
	"StudentInformation_SecondStudentsInformation_SecondStudentsName_Last",
	"StudentInformation_SecondStudentsInformation_DateOfBirth",
	"StudentInformation_ThirdStudentsInformation_ThirdStudentsName_First",
	"StudentInformation_ThirdStudentsInformation_ThirdStudentsName_Last",
	"StudentInformation_ThirdStudentsInformation_DateOfBirth",
	"StudentInformation_FourthStudentsInformation_FourthStudentsName_First",
	"StudentInformation_FourthStudentsInformation_FourthStudentsName_Last",
	"StudentInformation_FourthStudentsInformation_DateOfBirth",
}

func main() {
	if len(os.Args) > 1 {
		excelFileName = os.Args[1]
		sheetName = os.Args[2]
	} else {
		log.Fatal("Excel file name or sheet name are missing...")

		// For manual testing
		// "original-data-testing.xlsx" "Sheet1"
		// excelFileName = "original-data-testing.xlsx"
		// sheetName = "Sheet1"
	}

	f, err := excelize.OpenFile(excelFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	var labelIndexes []int

	columnLabels := rows[0]
	for index, label := range columnLabels {
		for _, value := range starLabels {
			if label == value {
				labelIndexes = append(labelIndexes, index)
			}
		}
	}

	fmt.Println(labelIndexes, len(labelIndexes))

	for _, row := range rows[1:] {
		father := fmt.Sprintf("%v %v", row[labelIndexes[0]], row[labelIndexes[1]])
		mother := fmt.Sprintf("%v %v", row[labelIndexes[2]], row[labelIndexes[3]])
		address := fmt.Sprintf("%v, %v, %v %v", row[labelIndexes[4]], row[labelIndexes[5]], row[labelIndexes[6]], row[labelIndexes[7]])
		phone := fmt.Sprintf("%v", row[labelIndexes[8]])
		email := fmt.Sprintf("%v", row[labelIndexes[9]])
		emergencyContact := fmt.Sprintf("%v %v", row[labelIndexes[10]], row[labelIndexes[11]])
		emergencyPhone := fmt.Sprintf("%v", row[labelIndexes[12]])
		student1 := fmt.Sprintf("%v %v", row[labelIndexes[13]], row[labelIndexes[14]])
		student1DOB := fmt.Sprintf("%v", row[labelIndexes[15]])
		student2 := fmt.Sprintf("%v %v", row[labelIndexes[16]], row[labelIndexes[17]])
		student2DOB := fmt.Sprintf("%v", row[labelIndexes[18]])
		student3 := fmt.Sprintf("%v %v", row[labelIndexes[19]], row[labelIndexes[20]])
		student3DOB := fmt.Sprintf("%v", row[labelIndexes[21]])
		student4 := fmt.Sprintf("%v %v", row[labelIndexes[22]], row[labelIndexes[23]])
		student4DOB := fmt.Sprintf("%v", row[labelIndexes[24]])

		cleanRegistrantData = append(cleanRegistrantData, registrantInfo{father, mother, address, phone, email, emergencyContact, emergencyPhone, student1, student1DOB, student2, student2DOB, student3, student3DOB, student4, student4DOB})
	}

	// Excel
	newf := excelize.NewFile()
	newf.SetSheetRow("Sheet1", "A1", &[]string{"Father's Name", "Mother's Name", "Address", "Phone Number", "Email", "Emergency Contact", "Emergency Phone #", "Student 1 Name", "Student 1 DOB", "Student 2 Name", "Student 2 DOB", "Student 3 Name", "Student 3 DOB", "Student 4 Name", "Student 4 DOB"})
	count := 2

	for _, data := range cleanRegistrantData {
		index, err := excelize.CoordinatesToCellName(1, count)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Adding %v to Excel sheet...\n", data.student1Name)
		newf.SetSheetRow("Sheet1", index, &[]string{data.fatherName, data.motherName, data.address, data.phoneNumber, data.email, data.emergencyContact, data.emergencyNumber, data.student1Name, data.student1DOB, data.student2Name, data.student2DOB, data.student3Name, data.student3DOB, data.student4Name, data.student4DOB})

		count++

		if err := newf.SaveAs(fmt.Sprintf("%v-students/a-clean-%v", excelFileName, excelFileName)); err != nil {
			fmt.Println(err)
		}

		// Word
		replaceMap := docx.PlaceholderMap{
			"FATHERS_NAME":      data.fatherName,
			"MOTHERS_NAME":      data.motherName,
			"ADDRESS":           data.address,
			"EMAIL":             data.email,
			"PHONE_NUMBER":      data.phoneNumber,
			"EMERGENCY_PHONE":   data.emergencyNumber,
			"EMERGENCY_CONTACT": data.emergencyContact,
			"CHILD_1":           data.student1Name,
			"CHILD_1_DOB":       data.student1DOB,
			"CHILD_1_TEAM":      "",
			"CHILD_2":           data.student2Name,
			"CHILD_2_DOB":       data.student2DOB,
			"CHILD_2_TEAM":      "",
			"CHILD_3":           data.student3Name,
			"CHILD_3_DOB":       data.student3DOB,
			"CHILD_3_TEAM":      "",
			"CHILD_4":           data.student4Name,
			"CHILD_4_DOB":       data.student4DOB,
			"CHILD_4_TEAM":      "",
		}

		doc, err := docx.Open("student-info-doc-format.docx")
		if err != nil {
			log.Fatal(err)
		}

		err = doc.ReplaceAll(replaceMap)
		if err != nil {
			log.Fatal(err)
		}

		docName := fmt.Sprintf("%v-students/%v-info.docx", excelFileName, data.student1Name)
		err = doc.WriteToFile(docName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created %v...\n", docName)
	}

}
