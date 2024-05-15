// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../src/IncredibleSummingServiceManager.sol" as incsqsm;
import {IncredibleSummingTaskManager} from "../src/IncredibleSummingTaskManager.sol";
import {BLSMockAVSDeployer} from "@eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract IncredibleSummingTaskManagerTest is BLSMockAVSDeployer {
    incsqsm.IncredibleSummingServiceManager sm;
    incsqsm.IncredibleSummingServiceManager smImplementation;
    IncredibleSummingTaskManager tm;
    IncredibleSummingTaskManager tmImplementation;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address aggregator =
        address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    address generator =
        address(uint160(uint256(keccak256(abi.encodePacked("generator")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        tmImplementation = new IncredibleSummingTaskManager(
            incsqsm.IRegistryCoordinator(address(registryCoordinator)),
            TASK_RESPONSE_WINDOW_BLOCK
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        tm = IncredibleSummingTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(tmImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        tm.initialize.selector,
                        pauserRegistry,
                        registryCoordinatorOwner,
                        aggregator,
                        generator
                    )
                )
            )
        );
    }

    function testCreateNewTask() public {
        bytes memory quorumNumbers = new bytes(0);
        uint64[3] memory array = [uint64(1), uint64(2), uint64(3)];
        cheats.prank(generator, generator);
        tm.createNewTask(array, 100, quorumNumbers);
        assertEq(tm.latestTaskNum(), 1);
    }
}
