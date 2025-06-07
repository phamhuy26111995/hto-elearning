export const TEACHER_ENDPOINT = `${import.meta.env.VITE_TEACHER_ENDPOINT}`;
export const API_URL = `${import.meta.env.VITE_API_URL}`;
export const ROUTES = {
    LOGIN: "/login",
    COURSE: "/course",
    STUDENT : "/student",
}

export const ROLES = [
    {
        name: 'Học sinh',
        value: 'STUDENT'
    },
    {
        name: 'Giáo viên',
        value: 'TEACHER'
    },
    {
        name: 'Quản lý',
        value: 'ADMIN'
    }
]