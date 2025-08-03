
export const ENDPOINT = `${import.meta.env.VITE_ENDPOINT}`
export const COMMON_ENDPOINT = "api/v1/common" as const



export const API_URL = `${import.meta.env.VITE_API_URL}`;
export const DEFAULT_TIMEOUT = 2000; // Default timeout for API requests in milliseconds
export const ROUTES = {
    LOGIN: "/login",
    COURSE: "/course",
    USER : "/user",
}

export const ROLES = [
    {
        name: 'Học sinh',
        value: 'STUDENT'
    },
    {
        name: 'Giáo viên',
        value: 'TEACHER'
    }
]