<template id="salz-game-view">
  <div id="salz-game-inner-view">
    <GameMenu v-if="!hideUI" />
    <FrameControl v-if="!hideUI" />
    <Hotkeys :viewport="viewport" />
    <Ranking v-if="showRanking" />
    <Help v-if="showHelp" />
    <Button label="Iterate" @click="iterate" class="iterate">Iterate</Button>
  </div>
</template>

<style lang="scss">
@import '~assets/css/colors';
#salz-game-inner-view {
  width: 100vw;
  height: 100vh;
  top: 0;
}

.cellInfoContainer {
  padding: 20px;
  position: fixed;
}

.iterate {
  background: red !important;
  position: absolute;
  top: 50px;
  z-index: 9999;
}
</style>

<script charset="utf-8">
import { mapState } from 'vuex';
import { Application } from 'pixi.js';
import { Viewport } from 'pixi-viewport';

// Game Rendering
import GameFrame from '@/lib/game/rendering/gameFrame';
import Frame from '@/lib/game/rendering/frame';
import Cell from '@/lib/game/entities/cell';
import { Color } from '@/lib/game/colors.js';
import { sketchyMedoid } from '@/lib/game/mathstuff';
import { EventBus } from '@/lib/eventBus';

import GameMenu from '@/components/Game/GameMenu';
import FrameControl from '@/components/Game/FrameControl';
import Hotkeys from '@/components/Game/Hotkeys';
import Ranking from '@/components/Game/Ranking';
import Help from '@/components/Game/Help';
import Button from '@/components/form/Button';

import Grid from '@/lib/game/conway';

export default {
  components: {
    GameMenu,
    FrameControl,
    Hotkeys,
    Ranking,
    Help,
    Button,
  },
  data() {
    return {
      turn: 0,
      grid: new Grid(0, 0),
      viewport: null,
      cellcolors: [
        Color.pink,
        Color.red,
        Color.orange,
        Color.yellow,
        Color.green,
        Color.turquiose,
        Color.cyan,
        Color.blue,
        Color.purple,
      ],
    };
  },
  computed: {
    ...mapState({
      user: state => state.login.username,
      userid: state => state.login.id,
      currentFrameNumber: state => state.game.activeFrame,
      totalFrames: state => state.game.framesLength,
      hideUI: state => state.game.hideUI,
      showRanking: state => state.game.showRanking,
      showHelp: state => state.game.showHelp,
    }),
    snapshots: () => require('@/dummy-data/snapshots').default,
    moves: () => require('@/dummy-data/moves').default,
  },
  // asyncData({ $axios }) {
  //   return $axios
  //     .$get('/frames')
  //     .then((res) => {
  //       return { index: res };
  //     })
  //     .catch((e) => {
  //       // eslint-disable-next-line
  //       console.warn(e);
  //       return { index: { frames: [] } };
  //     });
  // },
  beforeMount() {
    this.$store.dispatch('login/grabToken');
  },
  async mounted() {
    const wrapper = document.querySelector('#salz-game-inner-view');

    const app = new Application({
      backgroundColor: Color.black,
      resizeTo: wrapper,
      resolution: window.devicePixelRatio || 1,
    });
    wrapper.appendChild(app.view);
    // this.gameViewport = app.view;

    const vpWorldWidth = 1000;
    const vpWorldHeight = 1000;

    this.viewport = new Viewport({
      screenWidth: window.innerWidth,
      screenHeight: window.innerHeight,
      worldWidth: vpWorldWidth,
      worldHeight: vpWorldHeight,
      interaction: app.renderer.plugins.interaction,
    });

    // Set the centre to be the mediod of player's cells
    // const playerID = parseInt(this.userid, 10);

    this.viewport.moveCenter(0, 0);
    app.stage.addChild(this.viewport);
    this.viewport
      .drag()
      .pinch()
      .wheel()
      .decelerate()
      .clampZoom({
        minWidth: vpWorldWidth / 5,
        minHeight: vpWorldHeight / 5,
        maxWidth: vpWorldWidth * 5,
        maxHeight: vpWorldHeight * 5,
      });
    this.grid = await new Grid(50, 50);
    await this.grid.mountSnapshot(this.snapshots[0]);

    this.grid.cells.forEach(cell => {
      this.viewport.addChild(new Cell(cell.owner, cell.x, cell.y, this.cellcolors[0]));
    });
    this.turn = this.snapshots[0].turnid;

    // Add a helper function for viewport
    /**
     * Move viewport camera by dx and dy
     *
     * @param   {number}  dx  Pixels to move x by
     * @param   {number}  dy  Pixels to move y by
     */
    this.viewport.moveViewport = function(dx, dy) {
      this.viewport.moveCenter(this.viewport.center.x + dx, this.viewport.center.y + dy);
    };

    // EventBus.$on('updateFrameIndex', () => {
    //   gameFrame.mountFrame(frames[this.currentFrameNumber]);
    // });
  },
  methods: {
    iterate() {
      const nextMoves = this.moves.filter(moves => moves.turnid === this.turn);
      this.turn += 1;
      if (nextMoves) {
        this.grid.next(nextMoves[0].moves);
        this.updateCanvas();
      }
    },
    updateCanvas() {
      this.viewport.removeChildren();
      this.grid.cells.forEach(cell => {
        this.viewport.addChild(new Cell(cell.owner, cell.x, cell.y, this.cellcolors[0]));
      });
    },
  },
};
</script>
