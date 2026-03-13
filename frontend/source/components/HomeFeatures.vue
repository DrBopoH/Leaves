<template>
    <div class="features">
        <div class="feature-card">
            <div class="icon">
                <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"></polygon></svg>
            </div>
            <h3>Lightning Fast</h3>
            <p>Zero lag and instant message delivery, powered by our custom backend.</p>
        </div>

        <div class="feature-card">
            <div class="icon">
                <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect><path d="M7 11V7a5 5 0 0 1 10 0v4"></path></svg>
            </div>
            <h3>Secure by Design</h3>
            <p>Your data belongs to you. Absolute privacy in every single message.</p>
        </div>

        <div class="feature-card">
            <div class="icon">
                <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2v20M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path></svg>
            </div>
            <h3>Clean & Minimal</h3>
            <p>No ads, no unnecessary clutter. Just you and your conversations.</p>
        </div>
    </div>
</template>

<style scoped>
@keyframes fadeInUp {
    from { opacity: 0; transform: translateY(30px); }
    to { opacity: 1; transform: translateY(0); }
}

@keyframes float {
    0% { transform: translateY(0px); }
    50% { transform: translateY(-10px); }
    100% { transform: translateY(0px); }
}

.features {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: 30px;
    width: 100%;
    animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) 0.2s forwards;
    opacity: 0;
}

.feature-card {
    background: var(--color-surface);
    border: 1px solid var(--color-border);
    padding: 30px;
    border-radius: 16px;
    text-align: left;
    transition: all 0.3s ease;
    width: 100%;
}

.feature-card:hover {
    border-color: var(--color-accent-hover);
    transform: translateY(-5px);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
}

.feature-card .icon {
    font-size: 2.5rem;
    margin-bottom: 20px;
    animation: float 4s ease-in-out infinite;
}

.feature-card .icon svg {
    width: 32px;
    height: 32px;
    min-width: 32px;
    flex-shrink: 0;
    stroke: var(--color-accent);
    transition: stroke 0.3s ease;
}

.feature-card h3 {
    color: var(--color-text-primary);
    font-size: 1.3rem;
    margin-bottom: 15px;
    transition: color 0.3s ease;
}

.feature-card p {
    color: var(--color-text-secondary);
    line-height: 1.5;
    font-size: 0.95rem;
    transition: color 0.3s ease;
}
</style>