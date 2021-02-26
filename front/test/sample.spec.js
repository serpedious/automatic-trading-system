import { mount } from '@vue/test-utils'
import AppLogo from '@/components/AppLogo.vue'

describe('AppLogo', () => {
  test('exist AppLogo component', () => {
    const wrapper = mount(AppLogo)
    expect(wrapper.exists()).toBeTruthy()
  })
})