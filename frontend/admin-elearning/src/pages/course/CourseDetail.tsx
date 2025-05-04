import React from 'react'

import { useParams } from "react-router";

export default function CourseDetail() {
    const { courseId } = useParams();

  return (
    <div>CourseDetail {courseId}</div>
  )
}
