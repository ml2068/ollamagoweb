var converter = new showdown.Converter();
let conversationHistory = [];
let currentConversationId = 0;
const MAX_CONVERSATIONS = 3;
const MAX_CONVERSATION_LENGTH = 6656;

function send(e){
    e.preventDefault();
    var prompt = $("#prompt").val().trimEnd();
    var format_prompt=hljs.highlightAuto(prompt).value;
    var inputId = "input-" + uuidv4();
    $("#prompt").val("");
    autosize.update($("#prompt"));
    $("#printout").append(
        "<div class='px-3 py-3'>" +
        "<div id='" + inputId +
        "'style='white-space: pre-wrap;'>" +
        "</div>" +
        "</div>" +
        "<div class='prompt-message'>" + 
        "<div style='white-space: pre-wrap;'>" +
        "<h4>Question:</h4>"+
        format_prompt  +
        "</div>" +
        "<span class='message-loader js-loading spinner-border'></span>" +
        "</div>" +
        "\n"
    );        
    window.scrollTo({top: document.body.scrollHeight, behavior:'smooth' });
    runScript(prompt);          
    $(".js-logo").addClass("active");
};

function uuidv4() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
    var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
    return v.toString(16);
  });
}

$(document).ready(function(){  
    $('#prompt').keypress(function(event){        
        var keycode = (event.keyCode ? event.keyCode : event.which);
        if((keycode == 10 || keycode == 13) && event.ctrlKey){
            send(event);
            return false;
        }
    });       
    autosize($('#prompt'));    
});  

async function runScript(prompt, inputId) {    
    var outputId = "result-" + uuidv4();
    $("#printout").append(
        "<div class='px-3 py-3'>" + 
        "<div id='" + outputId + 
        "' style='white-space: pre-wrap;'>" +         
        "</div>" +
        "</div>" +
        "\n"
    );  

    // 将三轮对话内容插入#prompt内容
    var conversationText = "";
    conversationHistory.forEach(function(conversation) {
    conversationText += conversation.inputContent + "\n" + conversation.outputContent + "\n";
    });
    //var newPrompt = `请根据上述对话内容及问答的先后顺序:${conversationText}\n 回答问题：${prompt}，\n如果上下文有关就回答，不相关就直接按问题内容回答`;
    var newPrompt = `Please answer based on the conversation context and the order of the questions:\n ${conversationText}\n Answer the question: ${prompt},\n If relevant to the context, respond accordingly; otherwise, answer based on the question content only.`

    response = await fetch("/run", {
        method: "POST",
        headers: { "Content-Type": "application/json"},
        body: JSON.stringify({input: newPrompt}),
    });

    decoder = new TextDecoder();
    reader = response.body.getReader();
    while (true) {
        const { done, value } = await reader.read();
        if (done) break;
        $("#"+ outputId).append(decoder.decode(value));
        window.scrollTo({top: document.body.scrollHeight, behavior:'smooth' });
    }
    $("#printout").find("#"+ outputId).each(function() {
        $(this).attr('style', '');
    });
    $(".js-loading").removeClass("spinner-border");        
    $("#"+outputId).html(converter.makeHtml($("#"+outputId).html()));
    window.scrollTo({top: document.body.scrollHeight, behavior:'smooth' });
    hljs.highlightAll();
    
  // 保存对话历史
  var conversation = {
    id: currentConversationId,
    inputId: inputId,
    outputId: outputId,
    inputContent: prompt,
    outputContent: $("#" + outputId).html()
  };

  // 对话长度限定在4096字符
  if (conversation.inputContent.length + conversation.outputContent.length > MAX_CONVERSATION_LENGTH) {
    conversation.inputContent = conversation.inputContent.substring(0, MAX_CONVERSATION_LENGTH / 4);
    conversation.outputContent = conversation.outputContent.substring(0, MAX_CONVERSATION_LENGTH / 4);
  }
  
  conversationHistory.push(conversation);
  currentConversationId++;

  // 只存储三轮对话
  if (conversationHistory.length > MAX_CONVERSATIONS) {
    conversationHistory.shift();
  }
}

document.getElementById("btnSave").addEventListener("click", () => {
    let date = new Date();
    let fileName = `${date.getFullYear()}${date.getMonth()+1}${date.getDate()}${date.getHours()}${date.getMinutes()}${date.getSeconds()}`.replace(/\s/g, '')+Math.random().toString(36).substring(2,5);
    const txt = document.getElementById('printout').innerHTML+`-----(`+document.getElementById('llmtag').innerText+`)-----`;
    const headHtml = `<head lang="en">
<meta charset="UTF-8" />
<meta name="viewport" content="width=device-width, initial-scale=1" />
</head>`;
    const blob = new Blob([headHtml+txt], {type: "text/html"});
    const url = URL.createObjectURL(blob);
    const ele = document.createElement("A");
    ele.href = url;
    ele.download = `llm${fileName}.html`;
    ele.click();
    setTimeout(() => URL.revokeObjectURL(url), 1000);
});
