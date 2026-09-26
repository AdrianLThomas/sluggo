const gameFrame = document.getElementById("game-frame");
const keyPressDuration = 100;

function sendKey(key, type) {
    const canvas = gameFrame.contentDocument?.querySelector("canvas");
    const frameWindow = gameFrame.contentWindow;
    if (!canvas || !frameWindow) {
        return;
    }

    canvas.dispatchEvent(new frameWindow.KeyboardEvent(type, {
        bubbles: true,
        cancelable: true,
        code: key,
        key,
    }));
}

document.querySelectorAll(".control").forEach((control) => {
    control.addEventListener("click", () => {
        const key = control.dataset.key;
        sendKey(key, "keydown");
        window.setTimeout(() => sendKey(key, "keyup"), keyPressDuration);
    });
});
