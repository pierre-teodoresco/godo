export function formatDate(date: Date, format: string): string {
    return new Date(date).toLocaleString(format, { dateStyle: 'long', timeStyle: 'short' })
}