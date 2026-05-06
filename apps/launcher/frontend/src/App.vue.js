import { ref } from "vue";
import logo from "./assets/images/logo-universal.png";
import { Greet } from "../wailsjs/go/main/App.js";
const resultText = ref("Please enter your name below 👇");
const name = ref("");
function greet() {
    Greet(name.value).then((result) => (resultText.value = result));
}
debugger; /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {};
let __VLS_components;
let __VLS_directives;
// CSS variable injection 
// CSS variable injection end 
__VLS_asFunctionalElement(__VLS_intrinsicElements.main, __VLS_intrinsicElements.main)({});
__VLS_asFunctionalElement(__VLS_intrinsicElements.img)({
    alt: "Wails logo",
    id: "logo",
    src: (__VLS_ctx.logo),
});
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
    ...{ class: "result" },
    id: "result",
});
(__VLS_ctx.resultText);
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
    ...{ class: "input-box" },
    id: "input",
});
__VLS_asFunctionalElement(__VLS_intrinsicElements.input)({
    autocomplete: "off",
    value: (__VLS_ctx.name),
    ...{ class: "input" },
    id: "name",
    type: "text",
});
__VLS_asFunctionalElement(__VLS_intrinsicElements.button, __VLS_intrinsicElements.button)({
    ...{ onClick: (__VLS_ctx.greet) },
    ...{ class: "btn" },
});
/** @type {__VLS_StyleScopedClasses['result']} */ ;
/** @type {__VLS_StyleScopedClasses['input-box']} */ ;
/** @type {__VLS_StyleScopedClasses['input']} */ ;
/** @type {__VLS_StyleScopedClasses['btn']} */ ;
var __VLS_dollars;
const __VLS_self = (await import('vue')).defineComponent({
    setup() {
        return {
            logo: logo,
            resultText: resultText,
            name: name,
            greet: greet,
        };
    },
});
export default (await import('vue')).defineComponent({
    setup() {
        return {};
    },
});
; /* PartiallyEnd: #4569/main.vue */
