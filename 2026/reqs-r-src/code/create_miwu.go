package demo

import (
    "net/http"
)

func CreateMIWU(req http.Request) http.Response {
    miwu, err := requestToData(req)
    if err != nil {
    	return errorToResponse(err)
    }
    if err = checkMIWU(miwu); err != nil {
    	return errorToResponse(err)
    }
    dbMIWU := miwuToDb(miwu)

    dbMIWU, err = insertMIWU(dbMIWU)
    if err != nil {
    	return errorToResponse(err)
    }
    return dbMIWUToResponse(dbMIWU)
}

