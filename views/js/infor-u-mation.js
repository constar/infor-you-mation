function getCardFlow(msgs) {
    var cardflows = document.getElementsByClassName("cardflow");
    for (var i = 0; i < msgs.length; i++) {
        var card = document.createElement("div");
        card.setAttribute("class", "card");
        cardflows[i%2].appendChild(card);
    }
    var cards = document.getElementsByClassName("card");
    for (var i = 0; i < cards.length; i++) {
        var cardmain = document.createElement("div");
        cardmain.setAttribute("class", "cardmain");
        cards[i].appendChild(cardmain);
    }
    var cardmains = document.getElementsByClassName("cardmain");
    for (var i = 0; i < cardmains.length; i++) {
        var keyword = document.createElement("div");
        keyword.setAttribute("class", "keyword");
        cardmains[i].appendChild(keyword);
        var cardtitle = document.createElement("div");
        cardtitle.setAttribute("class", "cardtitle")
        cardmains[i].appendChild(cardtitle);
    }
    var keywords = document.getElementsByClassName("keyword");
    for (var i = 0; i < keywords.length; i++) {
        var divcenter = document.createElement("div");
        divcenter.setAttribute("class", "divcenter");
        keywords[i].appendChild(divcenter);
    }
    var divcenters = document.getElementsByClassName("divcenter");
    for (var i = 0; i < divcenters.length; i++) {
        var ak = document.createElement("a");
        ak.setAttribute("class", "ak");
        ak.innerHTML=msgs[i].Topic;
        ak.onclick = getDetails;
        divcenters[i].appendChild(ak);
        var add = document.createElement("P");
        add.setAttribute("class", "yesterdayadd");
        add.innerHTML="过去一天新增数：" + msgs[i].PastDayAdd;
        divcenters[i].appendChild(add);
    }
    var cardtitles = document.getElementsByClassName("cardtitle");
    for (var i = 0; i < cardtitles.length; i++) {
        var titlelist = document.createElement("ul");
        titlelist.setAttribute("class", "titlelist");
        cardtitles[i].appendChild(titlelist);
    }
    var titlelists = document.getElementsByClassName("titlelist");
    for (var i = 0; i < titlelists.length; i++) {
        for (var j = 0; j < 5; j++) {
            var singletitle = document.createElement("li");
            singletitle.setAttribute("class", "singletitle");
            titlelists[i].appendChild(singletitle);
        }
    }
    var singletitles = document.getElementsByClassName("singletitle");
    for (var i = 0; i < titlelists.length; i++) {
        for (var j = 0; j < 5; j++) {
            var al = document.createElement("a");
            al.setAttribute("class", "al");
            al.innerHTML=msgs[i].Feeds[j].Title;
            al.href=msgs[i].Feeds[j].Url;
            titlelists[i].childNodes[j].appendChild(al);
        }
    }
}
function getDetails() {
    var keyword = this.innerHTML;
    var scriptelement = document.createElement("script");
    var url = "http://api.ydy.yanyiwu.com/carddetail" + "?k=" + keyword + "&callback=showdetails";
    scriptelement.setAttribute("src", url)
    var body = document.getElementsByTagName("body")[0];
    var divcontent = document.getElementsByClassName("content content0")[0];
    body.removeChild(divcontent);
    var mainpart = document.createElement("div");
    mainpart.setAttribute("class", "mainpart");
    body.appendChild(mainpart);
    body.appendChild(scriptelement);
    var toppart = document.createElement("div");
    toppart.setAttribute("class", "toppart");
    mainpart.appendChild(toppart);
    var detailpart = document.createElement("div");
    detailpart.setAttribute("class", "detailpart");
    mainpart.appendChild(detailpart);
    var cardtopic = document.createElement("p");
    cardtopic.setAttribute("class", "cardtopic");
    cardtopic.innerHTML=keyword;
    toppart.appendChild(cardtopic);
    var yesterdayadd = document.createElement("p");
    yesterdayadd.setAttribute("class", "yesterdayadd");
    toppart.appendChild(yesterdayadd);
    var uldetail = document.createElement("ul");
    uldetail.setAttribute("class", "ul_detail");
    detailpart.appendChild(uldetail);
}
function showdetails(detail) {
    for (var i = 0; i < detail.Feeds.length; i++) {
        var lidetail = document.createElement("li");
        lidetail.setAttribute("class", "li_detail");
        var adetail = document.createElement("a");
        adetail.setAttribute("class", "a_detail");
        adetail.innerHTML=detail.Feeds[i].Title;
        adetail.href=detail.Feeds[i].Url;
        lidetail.appendChild(adetail);
        var uldetail = document.getElementsByClassName("ul_detail")[0];
        uldetail.appendChild(lidetail);
    }
    var yesterdayadd = document.getElementsByClassName("yesterdayadd")[0];
    yesterdayadd.innerHTML="过去一天新增数: "+ detail.PastDayAdd;
}
