function formatDate(inputDate) {
  var parts = inputDate.split("/");
  var dateObject = new Date(parts[2], parts[1] - 1, parts[0]);

  var formattedDate = Utilities.formatDate(dateObject, "GMT", "yyyy-MM-dd");

  return formattedDate;
}

function onFormSubmit(e) {
  requestData = e.namedValues
  var transformedData = {
    "email": requestData["Email"][0],
    "end_day": formatDate(requestData["End day"][0]),
    "first_name": requestData["First name"][0],
    "last_name": requestData["Last name"][0],
    "phone_number": requestData["Phone number"][0],
    "registration_time": requestData["Time"][0],
    "room_id": requestData["Room:"][0].split(',')[0],
    "start_day": formatDate(requestData["Start day"][0]),
    "student_id": requestData["Student ID "][0],
    "supervisor": requestData["Supervisor"][1]
  };
  Logger.log(JSON.stringify(transformedData))
  var data = JSON.stringify(transformedData)

  var api = "API"
  
  var options = {
    "method": "post",
    "contentType": "application/json",
    "payload": data
  };
  var response = UrlFetchApp.fetch(api, options);
  Logger.log(response.getContentText())
}
