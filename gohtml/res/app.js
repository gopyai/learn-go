// alert('Height: ' + window.screen.height + ', width: ' + window.screen.width);

let setOverlayPosition = () => {};
let isSmallMedia = () => {};
let showNav = () => {};
let hideNav = () => {};

//
// JS Identifier
//

class JSId {
    constructor() {this.id = 0;}

    generate() {
        this.id++;
        return "js" + this.id;
    }
}

let jsid = new JSId();

//

const reAttr = /[^a-zA-Z0-9\ _()!?,:\-&;=]/g;

function escapeAttr(s) {
    s = s.replace("&", "&amp;")
        .replace("<", "&lt;")
        .replace(">", "&gt;");
    return s.replace(reAttr, " ");
}

//

function autoHeight(el) {
    el.style.height = "5px";
    el.style.height = (el.scrollHeight + 20) + "px";
}

//

function htmlLoader(url, $el = $("main")) {
    const p = new Promise((resolve, reject) => {
        console.log("Get HTML:", url);
        $.ajax({
            type: "GET",
            url: url,
            dataType: "html",
            success: (html) => {
                $el.html(html);
                setOverlayPosition();
                resolve(html);
            },
            error: (jqxhr, textStatus) => {
                alert(`Load HTML error: ${url}`);
                reject(textStatus);
            },
        });
    });
    return p;
}

function apiCaller(method, url, data = {}, timeout = 10000) {
    const p = new Promise((resolve, reject) => {
        console.log(method, url, data);
        $.ajax({
            type: method,
            url: url,
            data: JSON.stringify(data),
            contentType: "application/json",
            dataType: "json",
            timeout,
            success: (data) => {
                console.log("Success:", method, url, data);
                resolve(data);
            },
            error: (jqxhr, textStatus) => {
                console.log("Error:", textStatus);
                alert(`Call api error: ${url}`);
                reject(textStatus);
            },
        });
    });
    return p;
}

//

function attachAjaxClick($el, p) {

    function f() {
        // console.log("Off handler ...");
        $el.off();
        p().finally(() => {
            // console.log("On handler ...");
            $el.click(f); // Re-attach click handler
        });
    }

    $el.click(f); // Attach click handler
}

//
// Menu and button handler
//

function handleMenu(name, url) {
    $(document).ready(() => {
        let $selected = $(".selected");
        let $button = $("#" + name);
        attachAjaxClick(
            $button,
            () => htmlLoader(url).then(html => {
                $selected.removeClass("selected");
                $button.addClass("selected");
                if (isSmallMedia()) hideNav();
            }).catch(reason => {}),
        );
    });
}

function handleButton(name, url) {
    $(document).ready(() => {
        attachAjaxClick(
            $("#" + name),
            () => htmlLoader(url).catch(reason => {}),
        );
    });
}

//
// Ajax functions
//

function handleAjaxButton(name, method, url, urlGetHTML = "", data = {}) {
    $(document).ready(() => {
        attachAjaxClick(
            $("#" + name),
            () => apiCaller(method, url, data).then(data => {
                if (urlGetHTML) {
                    htmlLoader(urlGetHTML).catch(reason => {});
                }
            }).catch(reason => {}),
        );
    });
}

function handleAjaxText(name, isGetHTML, urlCreate, saved = "", id = 0, urlUpdate = "", urlGetHTML = "") {
    if (!urlUpdate) urlUpdate = urlCreate;
    if (!urlGetHTML) urlGetHTML = urlCreate;
    $(document).ready(() => {
        let $saveButton = $("#" + name + "-save");
        let $cancelButton = $("#" + name + "-cancel");
        let $allButton = $("#" + name + "-button");
        let $inp = $("#" + name + "-input");
        let $loadingGif = $("#" + name + "-loading");

        let inpEdited = saved;
        let onFocusOut = () => {};

        $loadingGif.hide();
        $allButton.hide();
        $inp.val(saved); // Show saved value
        if ($inp[0].nodeName === "TEXTAREA") autoHeight($inp[0]);

        function showSavedValueAndEnableInput() {
            $loadingGif.hide(); // Hide loading gif
            $inp.val(saved); // Show saved value
            if ($inp[0].nodeName === "TEXTAREA") autoHeight($inp[0]);
            $inp.prop("disabled", false); // Enable input
        }

        function setGetHTMLHandler() {
            // Change input into button
            $inp.addClass("button");
            $inp.css("width", "100%");
            $inp.css("margin-right", "0");

            // Set input to read only and then remove all event handlers
            // and set click handler to get html for main section.
            $inp.attr("readonly", true);

            // Assign handler
            $inp.off(); // Remove all existing event handlers and assign a new one
            let url = `${urlGetHTML}/${id}`;

            attachAjaxClick($inp, () => htmlLoader(url).catch(reason => {}));
        }

        function setFocusHandler() {
            $inp.focusin(() => {
                onFocusOut = showSavedValueAndEnableInput;
                $inp.val(inpEdited); // Show edited value
                $allButton.show();
                if ($inp[0].nodeName === "TEXTAREA") autoHeight($inp[0]);
            });

            $inp.focusout(() => {
                $inp.prop("disabled", true); // Disable input
                inpEdited = $inp.val(); // Store input value to edited
                setTimeout(() => {
                    $allButton.hide();
                    onFocusOut(); // Callback
                }, 100);
            });
        }

        function setButtonHandler() {
            attachAjaxClick($saveButton, () => {
                onFocusOut = () => {};

                if (inpEdited === "") {
                    alert("Can not save empty string");
                    showSavedValueAndEnableInput();
                    return Promise.resolve();
                }
                $loadingGif.show(); // Show loading gif

                let method;
                let url;
                if (saved === "") {
                    method = "POST";
                    url = urlCreate;
                } else {
                    method = "PUT";
                    url = `${urlUpdate}/${id}`;
                }

                return apiCaller(method, url, {Value: inpEdited})
                    .then(data => {
                        saved = data.Value;
                        id = data.Id;
                        showSavedValueAndEnableInput();
                        if (isGetHTML) setGetHTMLHandler();
                    })
                    .catch(reason => {
                        showSavedValueAndEnableInput();
                    });
            });

            $cancelButton.click(() => {
                onFocusOut = () => {};
                if (saved === "") {
                    $("#" + name).remove();
                    return;
                }
                inpEdited = saved;
                showSavedValueAndEnableInput();
            });
        }

        if (isGetHTML && saved !== "") {
            setGetHTMLHandler();
        } else {
            setFocusHandler();
            setButtonHandler();
        }
    });
}

function handleAjaxSelect(selId, urlLoad, urlUpdate, fDataUpdate, fGetHTML) {
    let prev;
    $(document).ready(() => {
        let $sel = $(`#${selId}`);

        console.log("DEBUG:", urlLoad);

        // Load html options
        htmlLoader(urlLoad, $sel)
            .then(html => {
                prev = html;
            })
            .catch(reason => {
            });

        // Assign handler
        $sel.change(() => {
            $(`#${selId} option:selected`).each(function () {
                let id = $(this).val();
                let url = `${urlUpdate}/${id}`;
                let data = fDataUpdate(id);
                apiCaller("PUT", url, data)
                    .then(data => {
                        prev = fGetHTML(data);
                        $sel.html(prev);
                    })
                    .catch(reason => {
                        $sel.html(prev);
                    });
            });
        });
    });
}

function handleAjaxMultiText(items, isTextArea, isGetHTML, buttonId, divId, urlCreate, placeholder, urlUpdate = "", urlGetHTML = "") {

    function addItem(item) {
        // console.log("Add item:", item);
        let name = jsid.generate();

        let htmlInput;
        if (isTextArea) {
            htmlInput = "<textarea onkeyup=\"autoHeight(this)\" " +
                "placeholder=\"" + escapeAttr(placeholder) + "\" " +
                "id=\"" + name + "-input\"" +
                "></textarea>";
        }
        else {
            htmlInput = "<input type=\"text\" " +
                "placeholder=\"" + escapeAttr(placeholder) + "\" " +
                "id=\"" + name + "-input\"" +
                ">";
        }

        let html = "<form " +
            "id=\"" + name + "\"" +
            "><div class=\"col\"><div>" + htmlInput +
            "<img src=\"/res/loading.gif\" alt=\"\" height=\"20\" width=\"20\" " +
            "id=\"" + name + "-loading\"" +
            "></div><div " +
            "id=\"" + name + "-button\"" +
            "><span class=\"button\" " +
            "id=\"" + name + "-save\"" +
            ">Save</span><span class=\"button\" " +
            "id=\"" + name + "-cancel\"" +
            ">X</span></div></div></form>";
        $(`#${divId}`).append(html);
        handleAjaxText(name, isGetHTML, urlCreate, item.saved, item.id, urlUpdate, urlGetHTML);
    }

    $(document).ready(() => {
        // Add each item at start
        for (let v of items) addItem(v);

        // Add new item when "add" button is clicked
        $(`#${buttonId}`).click(() => {
            addItem({saved: ""});
        });
    });
}

//
// Main page interaction
//

$(document).ready(() => {
    // Overlay position
    let $estateEl = $("div.estate");
    let $overlayEl = $("div.overlay");
    setOverlayPosition = () => {
        $overlayEl.css("top", $estateEl.offset().top);
        $overlayEl.css("left", $estateEl.offset().left);
        $overlayEl.css("width", $estateEl.width());
        $overlayEl.css("height", $estateEl.height());
    };
    setOverlayPosition();
    $(window).resize(() => {
        setOverlayPosition();
    });

    //
    // Internal functions
    //

    let $actionButton = $("div.action");
    let $menuButton = $("header > div.left > i");
    let $listEl = $("div.list");
    let $navEl = $("nav");

    isSmallMedia = () => $navEl.css("position") === "absolute";
    showNav = () => {
        $navEl.show();
        $overlayEl.show();

    };
    hideNav = () => {
        $navEl.hide();
        $overlayEl.hide();
    };

    function toggleNavDisplay() {
        if ($navEl.is(":visible")) {
            hideNav();
        } else {
            showNav();
        }
    }

    //
    // Default style
    //

    function setDefaultStyle() {
        $listEl.hide();
        if (isSmallMedia()) {
            console.log("Small size media");
            hideNav();
        } else {
            console.log("Normal size media");
            showNav();
        }
    }

    setDefaultStyle();

    //
    // Handlers
    //

    $menuButton.click(() => {
        toggleNavDisplay();
    });

    $actionButton.click(() => {
        $actionButton.hide();
        $listEl.show();
    });

    $listEl.click(() => {
        $actionButton.show();
        $listEl.hide();
    });
});