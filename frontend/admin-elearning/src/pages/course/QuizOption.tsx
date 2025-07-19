import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { Input } from "@/components/ui/input";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { FormCourse } from "@/types/course";

import React, { useMemo } from "react";
import {
  Control,
  FormState,
  useFieldArray,
  useFormContext,
  UseFormGetValues,
  UseFormRegister,
  UseFormSetValue,
} from "react-hook-form";

type QuizOptionProps = {
  moduleIndex: number;
  questionIndex: number;
  quizIndex: number;
  register: UseFormRegister<FormCourse>;
  control: Control<FormCourse>;
  formState: FormState<FormCourse>;
  type: "SINGLE" | "MULTIPLE";
  setValue: UseFormSetValue<FormCourse>;
  getValues: UseFormGetValues<FormCourse>;
};

export default function QuizOption({
  questionIndex,
  quizIndex,
  moduleIndex,
  register,
  control,
  type,
  formState,
  setValue,
  getValues,
}: QuizOptionProps) {
  const {
    fields: optionFields,
    append: appendOption,
    remove: removeOption,
  } = useFieldArray<
    FormCourse,
    `modules.${number}.quizzes.${number}.questions.${number}.options`
  >({
    control,
    name: `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options`,
  });

  const [checkedIndex, setCheckedIndex] = React.useState(
    getValues(
      `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options`
    )?.findIndex((option: any) => option.isCorrect) || 0
  );
  
  console.log(checkedIndex);
  

  const correctAnswerPath = (index: number): any =>
    `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options.${index}.isCorrect`;

  const handleSingleChange = (selectedIndex: number) => {
    debugger
    optionFields.forEach((_, index) => {
      setValue(correctAnswerPath(index), index === selectedIndex);
      if(index === selectedIndex) setCheckedIndex(index);
    });
  };

  // const selectedSingleIndex = useMemo(() => {
  //   return optionFields.findIndex((field) => {
  //     return getValues(correctAnswerPath(optionFields.indexOf(field)));
  //   });
  // }, [optionFields, getValues]);


  return (
    <div className="space-y-4">
      {type === "SINGLE" ? (
        <RadioGroup
          key={optionFields.length}
          value={checkedIndex.toString()}
          onValueChange={(val) => handleSingleChange(Number(val))}
        >
          {optionFields.map((field, index) => (
            <div key={field.id} className="flex items-center gap-4">
              <RadioGroupItem  checked={getValues(correctAnswerPath(index))} value={index.toString()} />
              <Input
                className="w-full"
                placeholder={`Option ${index + 1}`}
                {...register(
                  `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options.${index}.optionContent`
                )}
              />
            </div>
          ))}
        </RadioGroup>
      ) : (
        optionFields.map((field, index) => (
          <div key={field.id} className="flex items-center gap-4">
            <Checkbox
              // @ts-ignore
              checked={getValues(correctAnswerPath(index))}
              onCheckedChange={(checked) =>
                // @ts-ignore
                setValue(correctAnswerPath(index), !!checked)
              }
            />
            <Input
              className="w-full"
              placeholder={`Option ${index + 1}`}
              {...register(
                `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options.${index}.optionContent`
              )}
            />
          </div>
        ))
      )}

      <Button
        type="button"
        onClick={() =>
          appendOption({
            optionContent: "",
            isCorrect: false,
            createdAt: "",
            updatedAt: "",
            optionId: 0,
            orderIndex: 0,
            questionId: 0,
          })
        }
      >
        Add Option
      </Button>
      <Button onClick={() => {
        console.log(getValues(
          `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options`
        ))
      }}>Get Value</Button>
    </div>
  );
}
