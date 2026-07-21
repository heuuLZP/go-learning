function copyText(text) {
  if (navigator.clipboard && window.isSecureContext) {
    return navigator.clipboard.writeText(text);
  }

  const textArea = document.createElement('textarea');
  textArea.value = text;
  textArea.setAttribute('readonly', '');
  textArea.style.position = 'fixed';
  textArea.style.opacity = '0';
  document.body.appendChild(textArea);
  textArea.select();

  try {
    const copied = document.execCommand('copy');
    document.body.removeChild(textArea);
    return copied ? Promise.resolve() : Promise.reject(new Error('Copy failed'));
  } catch (error) {
    document.body.removeChild(textArea);
    return Promise.reject(error);
  }
}

document.querySelectorAll('pre > code').forEach((code) => {
  const pre = code.parentElement;
  const wrapper = document.createElement('div');
  const button = document.createElement('button');

  wrapper.className = 'code-block';
  pre.parentNode.insertBefore(wrapper, pre);
  wrapper.appendChild(pre);

  button.className = 'copy-button';
  button.type = 'button';
  button.textContent = '复制';
  button.setAttribute('aria-label', '复制代码');
  wrapper.appendChild(button);

  button.addEventListener('click', async () => {
    try {
      await copyText(code.textContent);
      button.textContent = '已复制';
    } catch (error) {
      button.textContent = '复制失败';
    }

    window.setTimeout(() => {
      button.textContent = '复制';
    }, 1600);
  });
});
