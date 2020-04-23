import { TestData } from '../pages/testdata';
import { Panel } from '../pages/panel';
import { EditPanel } from '../pages/editPanel';
import { Graph } from '../pages/graph';
import { componentFactory } from '../support';

export const Components = {
  DataSource: {
    TestData,
  },
  Panels: {
    Panel,
    EditPanel,
    Visualization: {
      Graph,
    },
  },
  Drawer: {
    General: componentFactory({
      selectors: {
        title: (title: string) => `Panel inspector title ${title}`,
        expand: 'Panel inspector expand',
        contract: 'Panel inspector contract',
        close: 'Panel inspector close',
        rcContentWrapper: () => '.drawer-content-wrapper',
      },
    }),
  },
  PanelInspectData: componentFactory({
    selectors: {
      content: 'Panel inspector Data content',
    },
  }),
  PanelInspectStats: componentFactory({
    selectors: {
      content: 'Panel inspector Stats content',
    },
  }),
  PanelInspectJSON: componentFactory({
    selectors: {
      content: 'Panel inspector JSON content',
    },
  }),
  PanelInspectQuery: componentFactory({
    selectors: {
      content: 'Panel inspector Query content',
    },
  }),
  Tab: componentFactory({
    selectors: {
      title: (title: string) => `Tab ${title}`,
      active: () => '[class*="-activeTabStyle"]',
    },
  }),
  QueryEditorToolbarItem: componentFactory({
    selectors: {
      button: (title: string) => `QueryEditor toolbar item button ${title}`,
    },
  }),
  BackButton: componentFactory({
    selectors: {
      backArrow: 'Go Back button',
    },
  }),
};
