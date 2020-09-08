package w2g_http_package

import (
	"fmt"
	"net/http"
)

func FillAnswerHeader(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
}

func OptionsAnswer(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
}

func SetupHeaderResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func ResponseOK(w http.ResponseWriter, addedRecordString []byte) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	n, _ := fmt.Fprintf(w, string(addedRecordString))
	fmt.Println("Response was sent ", n, " bytes")
	return
}

func ResponseNoContent(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	n, _ := fmt.Fprintf(w, "")
	fmt.Println("Response was sent ", n, " bytes")
	return
}

func ResponseConflict(w http.ResponseWriter, msg string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusConflict)
	n, _ := fmt.Fprintf(w, "{\"message\":\""+msg+"\"}")
	fmt.Println("Response was sent ", n, " bytes")
	return
}

func ResponseBadRequest(w http.ResponseWriter, err error, message string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	errorString := "{\"error_message\":\"" + err.Error() + "\",\"message\":\"" + message + "\"}"
	http.Error(w, errorString, http.StatusBadRequest)
	return
}

func ResponseNotFound(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	n, _ := fmt.Fprintf(w, "")
	fmt.Println("Response was sent ", n, " bytes")
	return
}

func ResponseInternalServerError(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	n, _ := fmt.Fprintf(w, "")
	fmt.Println("Response was sent ", n, " bytes")
	return
}

func ResponseUnknown(w http.ResponseWriter, message string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	errorString := "{\"message\":\"" + message + "\"}"
	http.Error(w, errorString, http.StatusInternalServerError)
	return
}

func ResponseForbidden(w http.ResponseWriter, message string, frontEndAction string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	n, _ := fmt.Fprintf(w, "{\"message\":\""+message+"\", \"action\":\""+frontEndAction+"\"}")
	fmt.Println("Response was sent ", n, " bytes")
	return
}
