package n3json

import (
	"testing"
)

const jsonStr = `
{
  "array": [
	1, 2,"3",null,
	3, { "obj" : { "x" : "y", "subobj" : { "x1": null, "a1" : [1,2,3,"4"] } } },
	{"obj2":{"x2":"y2"},   "obj3" : {"x3" :  "y3"}, "attr4": null}
  ],  
  "boolean": true,      "null": null,
  "number": 123,
  "object": {
	"a": "b",
	"c": "d",
	"e": "f",
	"subobj" : {
		"g" : "h",
		"bool": true,
		"sub2obj" : {
			"g2": "h2",
			"bool"  :false
		}
	}
  },
  "string": "Hello   World"
}
`

const jsonStr1 = `
{  "Activity": {    "LearningResources": {      "LearningResourceRefId": "B7337698-BF6D-B193-7F79-A07B87211B93"    },    "-lang": "en",    "SoftwareRequirementList": {      "SoftwareRequirement": [        {          "SoftwareTitle": "Flash Player",          "Version": "9.0",          "Vendor": "Adobe"        },        {          "SoftwareTitle": "Python",          "Version": "3.0",          "OS": "Linux"        }      ]    },    "SourceObjects": {      "SourceObject": {        "#content": "A71ADBD3-D93D-A64B-7166-E420D50EDABC",        "-SIF_RefObject": "Lesson"      }    },    "Preamble": "This is a very funny comedy - students should have passing familiarity with Shakespeare",    "Evaluation": {      "-EvaluationType": "Inline",      "Description": "Students should be able to correctly identify all major characters."    },    "-RefId": "C27E1FCF-C163-485F-BEF0-F36F18A0493A",    "MaxAttemptsAllowed": "3",    "ActivityWeight": "5",    "ActivityTime": {      "FinishDate": "2002-09-12",      "DueDate": "2002-09-12",      "CreationDate": "2002-06-15",      "Duration": {        "#content": "30",        "-Units": "minute"      },      "StartDate": "2002-09-10"    },    "AssessmentRefId": "03EDB29E-8116-B450-0435-FA87E42A0AD2",    "Title": "Shakespeare Essay - Much Ado About Nothing",    "LearningStandards": {      "LearningStandardItemRefId": "9DB15CEA-B2C5-4F66-94C3-7D0A0CAEDDA4"    },    "Points": "50"  }}
`

const jsonStr2 = `[    {
        "id": "4947ED1F-1E94-4850-8B8F-35C653F51E9C",
        "actor":
        {
            "name": "Marjorie Amaya",            "mbox": "marjorie45@trashymail.com"
        },
        "verb":
        {
            "id": "http://adlnet.gov/expapi/verbs/completed",
            "display":
            {
                "en-US": "completed"
            }
        },
        "object":
        {
            "id": "http://example.com/assignments/English-8-1-A:1",
            "definition":
            {
                "type": "http://adlnet.gov/expapi/activities/assessment",
                "name": "English-8-1-A:1"
            }
        },
        "result":
        {
            "completion": "true",
            "success": "true",
            "score":
            {
                "scaled": 55,
                "min": 0,
                "max": 100
            },
            "duration": "PT209M"
        }
    },

    {
        "id": "0E97E476-8ED5-4601-A795-6A0CE3F0F40A",
        "actor":
        {
            "name": "Daniel Moffit",
            "mbox": "daniel46@pookmail.com"
        },
        "verb":
        {
            "id": "http://adlnet.gov/expapi/verbs/completed",
            "display":
            {
                "en-US": "completed"
            }
        },
        "object":
        {
            "id": "http://example.com/assignments/English-8-1-A:1",
            "definition":
            {
                "type": "http://adlnet.gov/expapi/activities/assessment",
                "name": "English-8-1-A:1"
            }
        },
        "result":
        {
            "completion": "true",
            "success": "true",
            "score":
            {
                "scaled": 57,
                "min": 0,
                "max": 100
            },
            "duration": "PT188M"
        }
    } ]`

func TestFormat(t *testing.T) {
	// jsonFmt := FmtFile("./json_fmt_test.json", 2)
	jsonFmt := Fmt(jsonStr2, 2)
	fPln(isJSON(jsonFmt))
	fPln(jsonFmt)
}
