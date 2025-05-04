import { apiService } from '@/api/apiService';
import useCourseStore from '@/store/course'
import React, { useEffect } from 'react'

export default function Course() {
  const { courses, setCourses } = useCourseStore();
  

  useEffect(() => {
    fetchCourse();
  }, [])
  

  async function fetchCourse() {
    const response = await apiService.get('/api/v1/teacher/courses');
    const { data } : any  = response;
    setCourses(data);
  }

  return (
    <div className='p-5'>


    </div>
  )
}
