import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import React from "react";

import { useParams } from "react-router";
import Module from "./Module";

export default function CourseDetail() {
  const { courseId } = useParams();

  return (
    <div className="flex flex-col p-7">
      <div>
        <h1>Create Course</h1>
      </div>

      <div>
        <div>Course title</div>
        <Input />
      </div>

      <div>
        <Module />
      </div>
    </div>
  );
}
