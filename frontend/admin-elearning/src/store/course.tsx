import { create } from 'zustand'

interface Course {
    courseId: number
    title: string
    description: string
}


interface CourseStore {
    courses: Course[]
    setCourses: (courses: Course[]) => void

}


const useCourseStore  = create<CourseStore>((set) => ({
     courses: [],
     setCourses: (courses: Course[]) => set({ courses })
}))

export default useCourseStore;