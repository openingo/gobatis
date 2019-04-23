/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package xml

import "encoding/xml"

type DynamicData struct {
    CharData string
    If       []If    `xml:"if"`
    Include  Include `xml:"include"`
    Set      Set     `xml:"set"`
    Where    Where   `xml:"where"`
}

type Select struct {
    XMLName       xml.Name
    Id            string `xml:"id,attr"`
    ParameterType string `xml:"parameterType,attr"`
    ParameterMap  string `xml:"parameterMap,attr"`
    ResultType    string `xml:"resultType,attr"`
    ResultMap     string `xml:"resultMap,attr"`
    FlushCache    string `xml:"flushCache,attr"`
    UseCache      string `xml:"useCache,attr"`
    Timeout       string `xml:"timeout,attr"`
    FetchSize     string `xml:"fetchSize,attr"`
    StatementType string `xml:"statementType,attr"`
    ResultSetType string `xml:"resultSetType,attr"`

    //If       []If    `xml:"if"`
    //Include Include `xml:"include"`
    //Where   Where   `xml:"where"`
    //Data    string  `xml:",chardata"`
    Data string `xml:",innerxml"`
}

type Insert struct {
    XMLName          xml.Name
    Id               string `xml:"id,attr"`
    ParameterType    string `xml:"parameterType,attr"`
    FlushCache       string `xml:"flushCache,attr"`
    Timeout          string `xml:"timeout,attr"`
    StatementType    string `xml:"statementType,attr"`
    UseGeneratedKeys string `xml:"useGeneratedKeys,attr"`
    KeyProperty      string `xml:"keyProperty,attr"`
    KeyColumn        string `xml:"keyColumn,attr"`

    //If       []If    `xml:"if"`
    //Include Include `xml:"include"`
    //Where   Where   `xml:"where"`
    //Data    string  `xml:",chardata"`
    Data string `xml:",innerxml"`
}

type Update struct {
    XMLName       xml.Name
    Id            string `xml:"id,attr"`
    ParameterType string `xml:"parameterType,attr"`
    FlushCache    string `xml:"flushCache,attr"`
    Timeout       string `xml:"timeout,attr"`
    StatementType string `xml:"statementType,attr"`

    //If       []If    `xml:"if"`
    //Include Include `xml:"include"`
    //Set     Set     `xml:"set"`
    //Where   Where   `xml:"where"`
    //Data    string  `xml:",chardata"`
    Data string `xml:",innerxml"`
}

type Delete struct {
    XMLName       xml.Name
    Id            string `xml:"id,attr"`
    ParameterType string `xml:"parameterType,attr"`
    FlushCache    string `xml:"flushCache,attr"`
    Timeout       string `xml:"timeout,attr"`
    StatementType string `xml:"statementType,attr"`

    //If       []If    `xml:"if"`
    //Include Include `xml:"include"`
    //Where   Where   `xml:"where"`
    //Data    string  `xml:",chardata"`
    Data string `xml:",innerxml"`
}
