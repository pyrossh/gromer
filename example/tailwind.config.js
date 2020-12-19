module.exports = {
    purge: {
        content: [
            './*.go',
        ],
        options: {
            safelist: ['flex', 'flex-row', 'justify-center', 'items-center'],
        }
    },
    theme: {},
    variants: {},
    plugins: [],
}