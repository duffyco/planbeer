package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type SessionCSVFileFormat struct {
	Sessions []SessionCSVFormat
}

type SessionCSVFormat struct {
	ZSessionID int
	ID int
	LogDate string
	StepName string
	TargetTemp float32
	WortTemp float32
	ThermoBlockTemp float32
	AmbientTemp float32
	DrainTemp float32
	ValvePosition int
	KegPumpOn bool
	DrainPumpOn bool
	ErrorCode int
	PauseReason int
	Rssi int
	NetSend int
	NetWait int
	NetRecv int
}

func Unmarshal(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}

	if( strings.HasPrefix( record[0], "#" ) ) {
		return nil;
	}


	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":			
			f.SetString(record[i])
		case "bool":
			bval, err := strconv.ParseBool(record[i])
			if err != nil {
				return err
			}
			f.SetBool(bval)
		case "int":
			ival, err := strconv.ParseInt(record[i], 10, 64)
			if err != nil {
				return err
			}
			f.SetInt( ival ) 
		case "float32":
			fval, err := strconv.ParseFloat(record[i], 10)
			if err != nil {
				return err
			}
			f.SetFloat(fval)
		default:
			return &UnsupportedType{f.Type().String()}
		}
	}
	return nil
}

func getImportableSessions( importPath string ) []SessionCSVFileFormat{
	csvfiles := ScanForRecipes( importPath )

	var sessionFiles []SessionCSVFileFormat

	//Load cvs
    for _, csvs := range csvfiles {
		sessionFile := importCSVSessionFile( importPath + "/" + csvs.Name() )

		if( SessionExists( sessionFile.Sessions[0].ZSessionID ) ) {
			continue;
		}

		sessionFiles = append( sessionFiles, importCSVSessionFile( importPath + "/" + csvs.Name() ) );
	}

	return sessionFiles;
}

func importCSVSessionFile(filename string) SessionCSVFileFormat {
    data, _ := ioutil.ReadFile(filename)

	var reader = csv.NewReader(bytes.NewReader( data ))
	reader.Comma = ','
	
	var importFile SessionCSVFileFormat
	var importLine SessionCSVFormat
	for {
		err := Unmarshal(reader, &importLine)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		
		importFile.Sessions = append( importFile.Sessions, importLine )
	}

	return importFile;
}

type FieldMismatch struct {
	expected, found int
}

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}

type UnsupportedType struct {
	Type string
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.Type
}
