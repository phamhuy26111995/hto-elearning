import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { FormCourse } from "@/types/course";
import React from "react";
import { useFieldArray, useFormContext } from "react-hook-form";
import QuizQuestion from "./QuizQuestion";
import { Button } from "@/components/ui/button";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";

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
      <Accordion
        type="single"
        collapsible
        className="w-full"
        defaultValue="item-1"
      >
        {quizFields.map((quiz, index) => (
          <AccordionItem key={quiz.id} value={quiz.id}>
            <AccordionTrigger
              onKeyDown={(e) => {
                const tagName = (e.target as HTMLElement).tagName.toLowerCase();
                const isTypingElement =
                  ["input", "textarea"].includes(tagName) ||
                  (e.target as HTMLElement).isContentEditable;

                if (isTypingElement && (e.key === " " || e.code === "Space")) {
                  // Đừng để Accordion xử lý nếu người dùng đang gõ trong input
                  e.stopPropagation();
                  return;
                }

                // Nếu không phải input thì để Accordion xử lý toggle như bình thường
              }}
            >
              <div className="flex w-full max-w-sm  items-center gap-3">
                <Label>Tiêu đề câu hỏi :</Label>
                <Input
                  // onClick={(e) => e.stopPropagation()}
                  // onMouseDown={(e) => e.stopPropagation()}
                  // onKeyDown={(e) => {

                  //   if (e.code === 'Space' || e.key === ' ') {
                  //     e.preventDefault();
                  //     e.stopPropagation();
                  //   }
                  // }}
                  {...register(`modules.${moduleIndex}.quizzes.${index}.title`)}
                />
              </div>
            </AccordionTrigger>
            <AccordionContent className="flex flex-col gap-4 text-balance">
              <QuizQuestion moduleIndex={moduleIndex} quizIndex={index} />
            </AccordionContent>
          </AccordionItem>
        ))}
      </Accordion>
      <Button
        onClick={() =>
          appendQuiz({
            title: "",
            questions: [],
            moduleId: moduleIndex,
            createdAt: "",
            updatedAt: "",
            orderIndex: 0,
            quizId: 0,
          })
        }
      >
        Add Quiz
      </Button>
    </div>
  );
}
