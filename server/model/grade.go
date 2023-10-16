package model

import (
	"encoding/xml"
	"time"
)

// Grades is a wrapper for a list of grades => XML stuff
type Grades struct {
	XMLName xml.Name `xml:"rowset"`
	Grades  []Grade  `xml:"row"`
}

type Grade struct {
	XMLName         xml.Name   `xml:"row"`
	Date            customDate `xml:"datum"`
	LectureNumber   string     `xml:"lv_nummer"`
	LectureSemester string     `xml:"lv_semester"`
	LectureTitle    string     `xml:"lv_titel"`
	Examiner        string     `xml:"pruefer_nachname"`
	Grade           string     `xml:"uninotenamekurz"`
	ExamType        string     `xml:"exam_typ_name"`
	Modus           string     `xml:"modus"`
	StudyID         string     `xml:"studienidentifikator"`
	StudyName       string     `xml:"studienbezeichnung"`
	ECTS            string     `xml:"lv_credits"`
}

type customDate struct {
	time.Time
}

func (c *customDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string

	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	if t, err := time.Parse("2006-01-02T15:04:05", v); err == nil {
		return err
	} else {
		c.Time = t
		return nil
	}
}

func (grade *Grade) CompareToEncrypted(encryptedGrade *EncryptedGrade) bool {
	return grade.LectureTitle == encryptedGrade.LectureTitle && grade.Grade == encryptedGrade.Grade
}

func (grade *Grade) ToEncryptedGrade(deviceId string) *EncryptedGrade {
	return &EncryptedGrade{
		Grade:        grade.Grade,
		DeviceID:     deviceId,
		LectureTitle: grade.LectureTitle,
	}
}
