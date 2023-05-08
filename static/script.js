function submitForm(event, url, handleInputId, valueInputId, responseElementId) {
    event.preventDefault();
    const handle = document.getElementById(handleInputId).value;
    const value = valueInputId ? document.getElementById(valueInputId).value : "";
  
    const xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4) {
        const responseElement = document.getElementById(responseElementId);
        responseElement.textContent = xhr.responseText;
        if (xhr.status === 200) {
          document.getElementById(handleInputId).value = "";
          if (valueInputId) document.getElementById(valueInputId).value = "";
        } else {
          document.getElementById("error").textContent = JSON.parse(xhr.responseText).error;
        }
      }
    };
    xhr.open("POST", url);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify({ handle, value }));
  }
  