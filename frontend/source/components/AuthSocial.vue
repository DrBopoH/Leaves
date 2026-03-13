<template>
    <div class="social-divider">
        <span>or continue with</span>
    </div>

    <div class="social-circles">
        <a href="#" class="s-circle">
            <img src="https://api.iconify.design/logos:google-icon.svg" class="social-icon" alt="Google" />
        </a>
        <a href="#" class="s-circle">
            <div class="social-icon github-icon" title="GitHub"></div>
        </a>
        <a href="#" class="s-circle">
            <div class="social-icon steam-gradient" title="Steam"></div>
        </a>
        <a href="#" class="s-circle">
            <img src="https://api.iconify.design/logos:discord-icon.svg" class="social-icon" alt="Discord" />
        </a>
    </div>
</template>

<style scoped>
.social-divider {
    text-align: center; margin-bottom: 16px; position: relative;
}
.social-divider::before, .social-divider::after {
    content: ''; position: absolute; top: 50%; width: 25%;
    height: 1px; background-color: var(--color-border, #0f1714);
}
.social-divider::before { left: 0; }
.social-divider::after { right: 0; }
.social-divider span {
    font-size: 12px; color: var(--color-text-secondary, #64615c);
    padding: 0 10px; background-color: var(--color-surface, #080b0a);
    transition: background-color 0.3s ease;
}
.social-circles {
    display: flex; justify-content: center; gap: 16px; margin-bottom: 24px;
}
.s-circle {
    width: 44px; height: 44px; background-color: var(--color-surface-alt, #040605);
    border: 1px solid var(--color-border, #0f1714); border-radius: 32px;
    display: flex; align-items: center; justify-content: center;
    cursor: pointer; transition: all 0.2s; text-decoration: none;
}
.s-circle:hover {
    background-color: var(--color-surface, #0f1714);
    border-color: var(--color-text-secondary, #64615c);
}
.social-icon { width: 24px; height: 24px; object-fit: contain; }
.github-icon {
    background-color: var(--color-text-primary, #c8c2b8);
    -webkit-mask: url('https://api.iconify.design/mdi:github.svg') no-repeat center / contain;
    mask: url('https://api.iconify.design/mdi:github.svg') no-repeat center / contain;
    transition: background-color 0.3s ease;
}
.steam-gradient {
    background: linear-gradient(to top, #00adee, #171a21);
    -webkit-mask: url('https://api.iconify.design/mdi:steam.svg') no-repeat center / contain;
    mask: url('https://api.iconify.design/mdi:steam.svg') no-repeat center / contain;
}
</style>