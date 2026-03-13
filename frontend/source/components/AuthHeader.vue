<template>
    <div class="brand-header">
        <div class="logo-wrapper">
            <img src="/Logo.svg" alt="Leaves" class="strict-logo" />
            <h1>Leaves</h1>
        </div>
        <p class="brand-subtitle">Create your account and start building your conversations today.</p>
    </div>
</template>

<style scoped>
.brand-header {
    text-align: center;
    margin-bottom: 24px; /* Отступ от логотипа до самой карточки */
}
.logo-wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 12px;
    margin-bottom: 12px;
}
.strict-logo { 
    width: 32px !important; 
    height: 32px !important; 
    min-width: 32px;
    object-fit: contain;
}
.logo-wrapper h1 {
    font-size: 28px; font-weight: 700; margin: 0;
    color: var(--color-text-primary, #c8c2b8);
    letter-spacing: -0.5px; transition: color 0.3s ease;
}
.brand-subtitle {
    font-size: 13px; color: var(--color-text-secondary, #8a867f);
    margin: 0; line-height: 1.5; transition: color 0.3s ease;
}
</style>