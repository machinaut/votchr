var votcher = "http://votchr.appspot.com/"

var insertVotchrImage = function(info, tab) {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", votcher+"votch", true);
    xhr.onreadystatechange = function() {
        if(xhr.readyState == 4) {
            if(xhr.status == 401) {
                console.log("User is not logged in");
                chrome.tabs.create({
                    url: votcher+"login"
                });
            } else {
                var votch = JSON.parse(xhr.responseText);
                console.log('Received votchr url: ' + votch['votch_url']);
                chrome.tabs.getSelected(null, function(tab){
                    if (tab === null || tab === undefined) {
                        console.log("Null tab selected");
                        return;
                    }
                    console.log("Got tab: " + tab.id);
                    chrome.tabs.sendRequest(tab.id, {
                        votch_url: votch['votch_url']
                    });
                });
            }
        }
    }
    xhr.send();
}

var id = chrome.contextMenus.create({
    "title": "Insert votchr image",
    "contexts": ["editable"],
    "onclick": insertVotchrImage
});
