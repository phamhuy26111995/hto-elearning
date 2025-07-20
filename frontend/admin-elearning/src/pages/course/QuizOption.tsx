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

import React, {
  forwardRef,
  useEffect,
  useImperativeHandle,
  useMemo,
} from "react";
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
  resetOptions: (callback?: any) => void;
};
const QuizOption = forwardRef(function QuizOption({
  questionIndex,
  quizIndex,
  moduleIndex,
  register,
  control,
  type,
  formState,
  setValue,
  getValues,
  resetOptions,
}: QuizOptionProps) {
  const {
    fields: optionFields,
    append: appendOption,
    remove: removeOption,
    replace,
  } = useFieldArray<
    FormCourse,
    `modules.${number}.quizzes.${number}.questions.${number}.options`
  >({
    control,
    name: `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options`,
  });

  useEffect(() => {
    if (resetOptions) {
      resetOptions(replace);
    }
  }, [resetOptions]);

  const [_, setTriggerRender] = React.useState(true);

  const correctAnswerPath = (index: number): any =>
    `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options.${index}.isCorrect`;

  const handleSingleChange = (selectedIndex: number) => {
    setTriggerRender((prev) => !prev);
    optionFields.forEach((_, index) => {
      setValue(correctAnswerPath(index), index === selectedIndex);
    });
  };

  const selectedOptions =
    type === "SINGLE"
      ? getValues(
          `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options`
        )?.findIndex((option: any) => option?.isCorrect) || 0
      : getValues(
          `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options`
        );

  return (
    <div className="space-y-4">
      {type === "SINGLE" ? (
        <RadioGroup
          value={selectedOptions!.toString()}
          onValueChange={(val) => handleSingleChange(Number(val))}
        >
          {optionFields.map((field, index) => (
            <div key={field.id} className="flex items-center gap-4">
              <RadioGroupItem value={index.toString()} />
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
              checked={getValues(correctAnswerPath(index))}
              onCheckedChange={(checked) => {
                setValue(correctAnswerPath(index), !!checked);
                setTriggerRender((prev) => !prev);
              }}
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
      <Button
        onClick={() => {
          console.log(
            getValues(
              `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options`
            )
          );
        }}
      >
        Get Value
      </Button>
    </div>
  );
});

export default QuizOption;
