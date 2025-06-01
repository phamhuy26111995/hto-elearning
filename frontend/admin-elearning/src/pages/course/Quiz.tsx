import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { FormCourse } from "@/types/course";
import React from "react";
import { useFieldArray, useFormContext } from "react-hook-form";
import QuizQuestion from "./QuizQuestion";

interface QuizProps {
  moduleIndex: number;
}

export default function Quiz({ moduleIndex }: QuizProps) {
  const {
    register,
    control,
    formState: { errors },
  } = useFormContext<FormCourse>();

  const {
    fields: quizFields,
    append: appendQuiz,
    remove: removeQuiz,
  } = useFieldArray<FormCourse, `modules.${number}.quizzes`>({
    control,
    name: `modules.${moduleIndex}.quizzes`,
  });

  return (
    <div>
      {quizFields.map((quiz, index) => (
        <>
          <div key={quiz.id}>
            <div className="grid w-full max-w-sm items-center gap-3">
              <Label>Tiêu đề câu hỏi :</Label>
              <Input
                {...register(`modules.${moduleIndex}.quizzes.${index}.title`)}
              />
            </div>

            <QuizQuestion moduleIndex={moduleIndex} quizIndex={index} />
          </div>
        </>
      ))}
    </div>
  );
}
