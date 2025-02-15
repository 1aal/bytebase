<template>
  <div
    v-if="task.taskCheckRunList.length > 0"
    class="flex items-start space-x-4"
  >
    <div class="textlabel h-[26px] inline-flex items-center">
      {{ $t("task.task-checks") }}
    </div>

    <TaskCheckBadgeBar
      :task-check-run-list="task.taskCheckRunList"
      @select-task-check-type="viewCheckRunDetail"
    />

    <RunTaskCheckButton v-if="allowRunTask" @run-checks="runChecks" />

    <BBModal
      v-if="state.showModal"
      :title="$t('task.check-result.title', { name: task.name })"
      class="!w-[56rem]"
      header-class="whitespace-pre-wrap break-all gap-x-1"
      @close="dismissDialog"
    >
      <div class="space-y-4">
        <div>
          <TaskCheckBadgeBar
            :task-check-run-list="task.taskCheckRunList"
            :allow-selection="true"
            :sticky-selection="true"
            :selected-task-check-type="state.selectedTaskCheckType"
            @select-task-check-type="viewCheckRunDetail"
          />
        </div>
        <BBTabFilter
          class="pt-4"
          :tab-item-list="tabItemList"
          :selected-index="state.selectedTabIndex"
          @select-index="
            (index: number) => {
              state.selectedTabIndex = index;
            }
          "
        />
        <TaskCheckRunPanel
          v-if="selectedTaskCheckRun"
          :task-check-run="selectedTaskCheckRun"
          :task="task"
        />
        <div class="pt-4 flex justify-end">
          <button
            type="button"
            class="btn-primary py-2 px-4"
            @click.prevent="dismissDialog"
          >
            {{ $t("common.close") }}
          </button>
        </div>
      </div>
    </BBModal>
  </div>
</template>

<script lang="ts">
import { cloneDeep } from "lodash-es";
import { computed, defineComponent, PropType, reactive } from "vue";
import { useI18n } from "vue-i18n";
import { BBTabFilterItem } from "@/bbkit/types";
import { Task, TaskCheckRun, TaskCheckStatus, TaskCheckType } from "@/types";
import { humanizeTs } from "@/utils";
import RunTaskCheckButton from "./RunTaskCheckButton.vue";
import TaskCheckBadgeBar from "./TaskCheckBadgeBar.vue";
import TaskCheckRunPanel from "./TaskCheckRunPanel.vue";

interface LocalState {
  showModal: boolean;
  selectedTaskCheckType: TaskCheckType | undefined;
  selectedTabIndex: number;
}

export default defineComponent({
  name: "TaskCheckBar",
  components: { TaskCheckBadgeBar, TaskCheckRunPanel, RunTaskCheckButton },
  props: {
    allowRunTask: {
      type: Boolean,
      default: true,
    },
    task: {
      required: true,
      type: Object as PropType<Task>,
    },
  },
  emits: ["run-checks"],
  setup(props, { emit }) {
    const { t } = useI18n();

    const state = reactive<LocalState>({
      showModal: false,
      selectedTaskCheckType: undefined,
      selectedTabIndex: 0,
    });

    const tabTaskCheckRunList = computed((): TaskCheckRun[] => {
      if (!state.selectedTaskCheckType) {
        return [];
      }

      const list: TaskCheckRun[] = [];
      for (const check of props.task.taskCheckRunList) {
        if (check.type == state.selectedTaskCheckType) {
          list.push(check);
        }
      }
      const clonedList = cloneDeep(list);
      clonedList.sort(
        (a: TaskCheckRun, b: TaskCheckRun) => b.createdTs - a.createdTs
      );
      return clonedList;
    });

    const tabItemList = computed((): BBTabFilterItem[] => {
      return tabTaskCheckRunList.value.map((item, index) => {
        return {
          title: index == 0 ? t("common.latest") : humanizeTs(item.createdTs),
          alert: false,
        };
      });
    });

    const selectedTaskCheckRun = computed(() => {
      const type = state.selectedTaskCheckType;
      const index = state.selectedTabIndex;
      if (!type) return undefined;
      return tabTaskCheckRunList.value[index];
    });

    // Returns the most severe status
    const taskCheckStatus = (taskCheckRun: TaskCheckRun): TaskCheckStatus => {
      let value: TaskCheckStatus = "SUCCESS";
      for (const result of taskCheckRun.result.resultList) {
        if (result.status == "ERROR") {
          return "ERROR";
        }
        if (result.status == "WARN") {
          value = "WARN";
        }
      }
      return value;
    };

    const viewCheckRunDetail = (type: TaskCheckType) => {
      state.selectedTaskCheckType = type;
      state.selectedTabIndex = 0;
      state.showModal = true;
    };

    const dismissDialog = () => {
      state.showModal = false;
      state.selectedTaskCheckType = undefined;
    };

    const runChecks = (taskList: Task[]) => {
      emit("run-checks", taskList);
    };

    return {
      state,
      tabTaskCheckRunList,
      tabItemList,
      selectedTaskCheckRun,
      taskCheckStatus,
      viewCheckRunDetail,
      dismissDialog,
      runChecks,
    };
  },
});
</script>
