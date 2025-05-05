/**
 * Formats a date into a human-readable format
 * @param d Date string to format
 * @returns Formatted date string
 */
export function formatDate(d: string): string {
    if (!d) return ''

    const date = new Date(d)
    const now = new Date()

    // Reset time parts to compare just the dates
    const dateWithoutTime = new Date(date.getFullYear(), date.getMonth(), date.getDate())
    const nowWithoutTime = new Date(now.getFullYear(), now.getMonth(), now.getDate())

    // Calculate diff in days
    const diffTime = nowWithoutTime.getTime() - dateWithoutTime.getTime()
    const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24))

    if (diffDays === 0) return 'Today'
    if (diffDays === 1) return 'Yesterday'
    if (diffDays > 1 && diffDays <= 7) return `${diffDays} days ago`

    return date.toLocaleDateString(undefined, { month: 'short', day: 'numeric', year: 'numeric' })
}

/**
 * Generates a consistent color based on a program name
 * @param name Program name
 * @returns Hex color code
 */
export function getProgramColor(name: string): string {
    const colors = [
        '#2080f0', '#18a058', '#f0a020', '#d03050',
        '#8a2be2', '#1890ff', '#52c41a', '#faad14',
        '#722ed1', '#eb2f96', '#f5222d', '#fa541c'
    ]
    let hash = 0
    for (let i = 0; i < name.length; i++) {
        hash = name.charCodeAt(i) + ((hash << 5) - hash)
    }
    hash = Math.abs(hash)
    return colors[hash % colors.length]
}

/**
 * Gets initials from a program name for avatar display
 * @param name Program name
 * @returns Up to 2 uppercase initials
 */
export function getInitials(name: string): string {
    if (!name) return ''
    return name
        .split(' ')
        .map(word => word[0])
        .join('')
        .toUpperCase()
        .substring(0, 2)
}