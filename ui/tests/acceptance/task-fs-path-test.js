import { currentURL } from '@ember/test-helpers';
import { Promise } from 'rsvp';
import { module, skip } from 'qunit';
import { setupApplicationTest, test } from 'ember-qunit';
import setupMirage from 'ember-cli-mirage/test-support/setup-mirage';
import Path from 'nomad-ui/tests/pages/allocations/task/fs/path';

let allocation;
let task;

module('Acceptance | task fs path', function(hooks) {
  setupApplicationTest(hooks);
  setupMirage(hooks);

  hooks.beforeEach(async function() {
    server.create('agent');
    server.create('node', 'forceIPv4');
    const job = server.create('job', { createAllocations: false });

    allocation = server.create('allocation', { jobId: job.id, clientStatus: 'running' });
    task = server.schema.taskStates.where({ allocationId: allocation.id }).models[0];
  });

  skip('visiting /allocations/:allocation_id/:task_name/fs/:path', async function(assert) {
    const paths = ['some-file.log', 'a/deep/path/to/a/file.log', '/', 'Unicode™®'];

    const testPath = async filePath => {
      await Path.visit({ id: allocation.id, name: task.name, path: filePath });
      assert.equal(
        currentURL(),
        `/allocations/${allocation.id}/${task.name}/fs/${encodeURIComponent(filePath)}`,
        'No redirect'
      );
      assert.ok(Path.tempTitle.includes(filePath), `Temp title includes path, ${filePath}`);
    };

    await paths.reduce(async (prev, filePath) => {
      await prev;
      return testPath(filePath);
    }, Promise.resolve());
  });

  test('visiting /allocations/:allocation_id/:task_name/fs/somewhere', async function(assert) {
    this.server.get('/client/fs/ls/:allocation_id', () => {
      return [{ Name: 'jorts' }, { Name: 'jants' }, { Name: 'directory', IsDir: true }];
    });

    await Path.visit({ id: allocation.id, name: task.name, path: 'somewhere' });

    assert.equal(Path.entries.length, 3);
    assert.equal(Path.entries[0].name, 'jorts');
    assert.ok(Path.entries[0].isFile);

    assert.equal(Path.entries[1].name, 'jants');

    assert.equal(Path.entries[2].name, 'directory');
    assert.ok(Path.entries[2].isDirectory);
  });
});
