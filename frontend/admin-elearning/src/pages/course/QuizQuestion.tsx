import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { FormCourse } from "@/types/course";
import React, { useRef } from "react";
import { useFieldArray, useFormContext } from "react-hook-form";
import QuizOption from "./QuizOption";
import useModalStore from "@/store/modal";
import { Card } from "@/components/ui/card";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

interface QuizQuestionProps {
  quizIndex: number;
  moduleIndex: number;
}

export default function QuizQuestion({
  quizIndex,
  moduleIndex,
}: QuizQuestionProps) {
  const { register, control, formState, getValues, setValue } =
    useFormContext<FormCourse>();

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
    <div className="flex flex-col gap-4">
      <Button
        type="button"
        onClick={() =>
          appendQuestion({
            questionId: 0,
            questionContent: "",
            questionType: "MULTIPLE",
            options: [],
            createdAt: "",
            updatedAt: "",

            orderIndex: questionFields.length,
            quizId: 0,
          })
        }
      >
        +
      </Button>
      {questionFields.map((question, index) => (
        <QuestionItem
          question={question}
          index={index}
          moduleIndex={moduleIndex}
          quizIndex={quizIndex}
          removeQuestion={removeQuestion}
        />
      ))}
    </div>
  );
}

function QuestionItem({
  question,
  index,
  moduleIndex,
  quizIndex,
  removeQuestion,
}: any) {
  const { register, control, formState, getValues, setValue } =
    useFormContext<FormCourse>();
  const openModal = useModalStore((state) => state.openModal);
  const [questionType, setQuestionType] = React.useState(question.questionType);
  const quizOptionRef = useRef<any>(null);


  function onChangeSelect(value: string) {
    setQuestionType(value);
    if(quizOptionRef.current){
      quizOptionRef.current();
    }
  }

 

  return (
    <Card className="p-8" key={question.id}>
      <div className="flex items-center gap-4">
        <div
          onClick={() => removeQuestion(index)}
          className="border-amber-300 rounded-2xl"
        >
          -
        </div>
        <div className="flex w-full max-w-sm items-center gap-3">
          <Label>Tên câu hỏi :</Label>
          <Input
            {...register(
              `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${index}.questionContent`
            )}
          />
        </div>
        <div className="flex w-full max-w-sm items-center gap-3">
          <Label>Loại câu hỏi</Label>
          <Select
            onValueChange={(value) => onChangeSelect(value)}
            defaultValue={question.questionType}
            {...register(
              `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${index}.questionType`
            )}
          >
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Select option" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="SINGLE">Single</SelectItem>
              <SelectItem value="MULTIPLE">Multiple</SelectItem>
            </SelectContent>
          </Select>
        </div>
        <Button
          onClick={() =>
            openModal(
              <QuizOption
                ref={quizOptionRef}
                key={"option_" + question.id}
                moduleIndex={moduleIndex}
                questionIndex={index}
                quizIndex={quizIndex}
                control={control}
                register={register}
                formState={formState}
                type={questionType}
                resetOptions={(callback : any) => quizOptionRef.current = callback}
                getValues={getValues}
                setValue={setValue}
              />
            )
          }
        >
          Edit Option
        </Button>
      </div>
    </Card>
  );
}
