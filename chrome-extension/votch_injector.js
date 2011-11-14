var insertHtmlAtCurrent = function(html) {
    console.log("In insert html");
    var selection = window.getSelection();
    if(selection.baseNode === null) {
        console.log("Selection is null");
        return;
    }
    var range = selection.getRangeAt(0);
    console.log("got range");
    var node = range.createContextualFragment(html);
    console.log("created node");
    range.insertNode(node);
    consol.log("inserted node");
}

chrome.extension.onRequest.addListener(
    function(request,sender){
        console.log(request);
        if(request['votch_url']){
            console.log("here");
            var html = '<img src="' + request['votch_url'] + '">';
            insertHtmlAtCurrent(html);
        }
    }
);

console.log("loaded votch_injector");
