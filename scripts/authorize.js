function Authorize(key, id) {
    // alert("victory shall be mine")
    // var myHeaders = new Headers()
    // myHeaders.append('Authorization', `Basic ${key}`);
    // myHeaders.append('Content-Type', 'text/plain');

    // var raw = '{\'authorized\': true}';

    // var requestOptions = {
    //     method: 'POST',
    //     headers: myHeaders,
    //     body: raw,
    //     redirect: 'follow'
    // };

    // fetch(`https://api.tailscale.com/api/v2/device/${id}/authorized`, requestOptions)
    // .then(response => response.text())
    // .then(result => console.log(result))
    // .catch(error => console.log('error', error));
    var myHeaders = new Headers();
    myHeaders.append("Authorization", "Basic OTczODk1Mjk0NzA4OTU4NA==");
    myHeaders.append("Content-Type", "text/plain");

    var raw = "{\"authorized\": true}";

    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };

    fetch("https://api.tailscale.com/api/v2/device/9738952947089584/authorized", requestOptions)
      .then(response => response.text())
      .then(result => console.log(result))
      .catch(error => console.log('error', error));
}