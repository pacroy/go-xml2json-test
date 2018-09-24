package main

import (
	"fmt"
	"strings"

	xj "github.com/basgys/goxml2json"
)

func main() {
	// xml is an io.Reader
	xml := strings.NewReader(`<?xml version="1.0" encoding="utf-8"?><cov:result name="ADT_ROOT_NODE" xmlns:cov="http://www.sap.com/adt/cov"><atom:link href="/sap/bc/adt/runtime/traces/coverage/results/0800279158BE1ED8AFF82D4FE1EF3922" rel="self" xmlns:atom="http://www.w3.org/2005/Atom"/><atom:link href="/sap/bc/adt/runtime/traces/coverage/results/0800279158BE1ED8AFF82D4FE1EF3922/statements" rel="http://www.sap.com/adt/relations/runtime/traces/coverage/results/bulkstatements" xmlns:atom="http://www.w3.org/2005/Atom"/><nodes><node><adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext/source/main#start=14,6" adtcore:type="CLAS/OCI" adtcore:name="ZCL_ZSADL_TEST_DPC_EXT========CP" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/><coverages><coverage type="branch" total="1" executed="0"/><coverage type="procedure" total="1" executed="0"/><coverage type="statement" total="3" executed="0"/></coverages><nodes><node><adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext/source/main#start=14,6" adtcore:type="CLAS/OCI" adtcore:name="ZCL_ZSADL_TEST_DPC_EXT" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/><coverages><coverage type="branch" total="1" executed="0"/><coverage type="procedure" total="1" executed="0"/><coverage type="statement" total="3" executed="0"/></coverages><atom:link href="/sap/bc/adt/runtime/traces/coverage/results/0800279158BE1ED8AFF82D4FE1EF3922/statements/ADT_ROOT_NODE/ZCL_ZSADL_TEST_DPC_EXT%3d%3d%3d%3d%3d%3d%3d%3dCP/ZCL_ZSADL_TEST_DPC_EXT" rel="http://www.sap.com/adt/relations/runtime/traces/coverage/results/statements" type="application/xml+scov" xmlns:atom="http://www.w3.org/2005/Atom"/><nodes><node><adtcore:objectReference adtcore:uri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext/source/main#start=17,9" adtcore:type="CLAS/OM" adtcore:name="T005TSET_UPDATE_ENTITY" adtcore:parentUri="/sap/bc/adt/oo/classes/zcl_zsadl_test_dpc_ext" adtcore:packageName="$SADL_TEST" xmlns:adtcore="http://www.sap.com/adt/core"/><coverages><coverage type="branch" total="1" executed="0"/><coverage type="procedure" total="1" executed="0"/><coverage type="statement" total="3" executed="0"/></coverages></node></nodes></node></nodes></node></nodes></cov:result>`)
	json, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	fmt.Println(json.String())
	// {"hello": "world"}
}
