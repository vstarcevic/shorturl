 function copyLink() {
    var copyText = document.getElementById("shortUrl");
    copyText.select();
    copyText.setSelectionRange(0, 99999); // For mobile devices
    navigator.clipboard.writeText(copyText.value);
}

async function shortenUrl() {
   hideResponseSections()

   if (!validate()) {
    showErrorSection()
    return;
   }
  
    request = { longUrl: document.getElementById("longUrl").value }
    var resp = await callShortUrlApi(request);
    if (resp != undefined && resp != '' && resp.length > 0) {
      document.getElementById("shortUrl").value = resp;
      showSuccessSection();
    }
}

const callShortUrlApi = async (obj) =>  {
  let options = {
      method: "POST",
      headers: {
          "Content-Type":"application/json",
      },
      body: JSON.stringify(obj)
  }
 
  try {
    const promise = await fetch('/shorten/', options);
    const response = await promise.json();

    if (response.error) {
      showErrorSection();
      return;
    }

    const shortUrl = response.data;
    return shortUrl;

  } catch (err) {
    showErrorSection();
  }
}

function hideResponseSections() {
  document.querySelector("#success-section").style.display = 'none'
  document.querySelector("#error-section").style.display = 'none'
}

function showErrorSection() {
  document.querySelector("#error-section").style.display = 'block'
}

function showSuccessSection() {
  document.querySelector("#success-section").style.display = 'block'
}

function validate() {
  return document.getElementById("longUrl").value.startsWith("http")
}