import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { FormCourse } from "@/types/course";
import React from "react";
import { useFieldArray, useFormContext } from "react-hook-form";
import QuizOption from "./QuizOption";

interface QuizQuestionProps {
  quizIndex: number;
  moduleIndex: number;
}

export default function QuizQuestion({
  quizIndex,
  moduleIndex,
}: QuizQuestionProps) {
  const {
    register,
    control,
    formState: { errors },
  } = useFormContext<FormCourse>();

  const {
    fields: questionFields,
    append: appendQuestion,
    remove: removeQuestion,
  } = useFieldArray<
    FormCourse,
    `modules.${number}.quizzes.${number}.questions`
  >({
    control,
    name: `modules.${moduleIndex}.quizzes.${quizIndex}.questions`,
  });

  return (
    <div>
      {questionFields.map((question, index) => (
        <div key={question.id}>
          <div className="grid w-full max-w-sm items-center gap-3">
            <Label>Tên câu hỏi :</Label>
            <Input
              {...register(
                `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${index}.questionContent`
              )}
            />
          </div>
          <div className="grid w-full max-w-sm items-center gap-3">
            <Label>Loại câu hỏi</Label>
            <Input
              {...register(
                `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${index}.questionType`
              )}
            />
          </div>
          <QuizOption key={"option_" + question.id} moduleIndex={moduleIndex} questionIndex={index} quizIndex={quizIndex} />
        </div>
      ))}

      <Button onClick={() => appendQuestion({
        questionId : 0,
        questionContent: "",
        questionType: "",
        createdAt : "",
        updatedAt : "",

        orderIndex : questionFields.length,
        quizId : 0
      })}>
        Add Question
      </Button>
    </div>
  );
}
