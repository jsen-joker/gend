import Vue from 'vue'
import Router from 'vue-router'
// import HelloWorld from '@/components/HelloWorld'
import Index from '@/components/index'
import Use from '@/components/use'
import QuickStart from '@/components/useCom/QuickStart'
import JokerBoot from '@/components/useCom/JokerBoot'
import JokerCore from '@/components/useCom/JokerCore'
import JokerProgram from '@/components/useCom/JokerProgram'
import JokerAnnotation from '@/components/useCom/JokerAnnotation'
import JokerPackage from '@/components/useCom/JokerPackage'
import JokerUse from '@/components/useCom/JokerUse'
import JokerUseJs from '@/components/useCom/JokerUse.js.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index
    },
    {
      path: '/use',
      name: 'Use',
      component: Use,
      children: [
        {
          path: 'quick-start',
          name: 'QuickStart',
          component: QuickStart
        },
        {
          path: 'joker-boot',
          name: 'JokerBoot',
          component: JokerBoot
        },
        {
          path: 'joker-core',
          name: 'JokerCore',
          component: JokerCore
        },
        {
          path: 'joker-program',
          name: 'JokerProgram',
          component: JokerProgram
        },
        {
          path: 'joker-annotation',
          name: 'JokerAnnotation',
          component: JokerAnnotation
        },
        {
          path: 'joker-package',
          name: 'JokerPackage',
          component: JokerPackage
        },
        {
          path: 'joker-use',
          name: 'JokerUse',
          component: JokerUse
        },
        {
          path: 'joker-use-js',
          name: 'JokerUseJs',
          component: JokerUseJs
        },
        {
          path: '',
          name: 'QuickStart',
          component: QuickStart
        }
      ]
    }
  ]
  // mode: 'history',
  // srcollBehavior (to, from, savedPosition) {
  //   if (to.hash) {
  //     console.log(to.hash)
  //     return {
  //       selector: to.hash
  //     }
  //   }
  // }
})
