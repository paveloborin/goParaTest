<?php
namespace functional;

class Simple7Test extends FunctionalTestCase
{
    public function testSomeActivitySimulation(): void
    {
        sleep(10);
        $this->assertTrue(
            true
        );
    }

}